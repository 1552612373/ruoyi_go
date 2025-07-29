package dao

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// SysMenu 菜单权限表
type SysMenu struct {
	// 菜单ID
	MenuID int64 `json:"menuId" gorm:"column:menu_id;primaryKey;autoIncrement"`

	// 菜单名称
	MenuName string `json:"menuName" gorm:"column:menu_name"`

	// 父菜单ID
	ParentID int64 `json:"parentId" gorm:"column:parent_id;default:0"`

	// 显示顺序
	OrderNum int `json:"orderNum" gorm:"column:order_num;default:0"`

	// 路由地址
	Path string `json:"path" gorm:"column:path;default:''"`

	// 组件路径
	Component string `json:"component" gorm:"column:component;default:null"`

	// 路由参数
	Query string `json:"query" gorm:"column:query;default:null"`

	// 路由名称
	RouteName string `json:"routeName" gorm:"column:route_name;default:''"`

	// 是否为外链（0是 1否）
	IsFrame int `json:"isFrame" gorm:"column:is_frame;default:1"`

	// 是否缓存（0缓存 1不缓存）
	IsCache int `json:"isCache" gorm:"column:is_cache;default:0"`

	// 菜单类型（M目录 C菜单 F按钮）
	MenuType string `json:"menuType" gorm:"column:menu_type;default:''"`

	// 菜单状态（0显示 1隐藏）
	Visible string `json:"visible" gorm:"column:visible;default:0"`

	// 菜单状态（0正常 1停用）
	Status string `json:"status" gorm:"column:status;default:0"`

	// 权限标识
	Perms string `json:"perms" gorm:"column:perms;default:null"`

	// 菜单图标
	Icon string `json:"icon" gorm:"column:icon;default:'#'"`

	// 创建者
	CreateBy string `json:"createBy" gorm:"column:create_by;default:''"`

	// 创建时间
	CreateTime time.Time `json:"createTime" gorm:"column:create_time"`

	// 更新者
	UpdateBy string `json:"updateBy" gorm:"column:update_by;default:''"`

	// 更新时间
	UpdateTime time.Time `json:"updateTime" gorm:"column:update_time"`

	// 备注
	Remark string `json:"remark" gorm:"column:remark;default:''"`
}

type MenuTreeNode struct {
	ID       int64           `json:"id"`
	Label    string          `json:"label"`
	Disabled bool            `json:"disabled"`
	Children []*MenuTreeNode `json:"children,omitempty"`
}

type SysMenuDAO struct {
	db *gorm.DB
}

func NewSysMenuDAO(db *gorm.DB) *SysMenuDAO {
	return &SysMenuDAO{
		db: db,
	}
}

func (dao *SysMenuDAO) Insert(ctx context.Context, obj SysMenu) error {
	err := dao.db.WithContext(ctx).Create(&obj).Error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflictsErrNo uint16 = 1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			// 唯一键冲突
			return errors.New("ZT唯一键冲突")
		}
	}
	return err
}

func (dao *SysMenuDAO) QueryList(ctx context.Context, pageNum int, pageSize int) ([]SysMenu, int, error) {
	objList := []SysMenu{}
	db := dao.db.WithContext(ctx).Model(&SysMenu{})

	var total int64

	// 查询总数
	db.Count(&total)

	// 分页处理
	if pageNum <= 0 {
		pageNum = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	// 执行分页查询
	err := db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&objList).Error

	return objList, int(total), err
}

func (dao *SysMenuDAO) Update(ctx context.Context, obj SysMenu) error {
	err := dao.db.WithContext(ctx).Model(&obj).Where("menu_id = ?", obj.MenuID).Updates(obj).Error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflictsErrNo uint16 = 1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			// 唯一键冲突
			return errors.New("ZT唯一键冲突")
		}
	}
	return err
}

func (dao *SysMenuDAO) QueryById(ctx context.Context, id int64) (SysMenu, error) {
	obj := SysMenu{}
	err := dao.db.WithContext(ctx).Where("menu_id = ?", id).First(&obj)
	return obj, err.Error
}

func (dao *SysMenuDAO) DeleteById(ctx context.Context, id int64) error {
	err := dao.db.WithContext(ctx).Where("menu_id = ?", id).Delete(&SysMenu{}).Error
	return err
}

// GetPermissionsByUserID 根据用户ID查询其拥有的所有权限字符串 (perms)
func (dao *SysMenuDAO) GetPermissionsByUserID(ctx context.Context, userID int64) ([]string, error) {
	var permissions []string

	// 使用 GORM 的 Joins 查询
	// 从 sys_menu 开始，通过 sys_role_menu 关联到 sys_role，再通过 sys_user_role 关联到 sys_user
	err := dao.db.WithContext(ctx).
		Table("sys_menu m").
		Joins("JOIN sys_role_menu rm ON m.menu_id = rm.menu_id").
		Joins("JOIN sys_role r ON rm.role_id = r.role_id").
		Joins("JOIN sys_user_role ur ON r.role_id = ur.role_id").
		Joins("JOIN sys_user u ON ur.user_id = u.user_id").
		Where("u.user_id = ?", userID).
		// 去重，因为一个用户可能通过多个角色拥有同一个权限
		Distinct("m.perms").
		Pluck("m.perms", &permissions).Error

	// 处理错误
	if err != nil {
		return nil, fmt.Errorf("failed to query permissions for user %d: %w", userID, err)
	}

	return permissions, nil
}

// GetRoleKeysByUserID 根据用户ID查询其拥有的所有角色标识 (role_key)
// 返回的 roleKeys 列表可能包含重复项，但通常不会，因为一个用户对同一个角色只会有一次关联。
func (dao *SysMenuDAO) GetRoleKeysByUserID(ctx context.Context, userID int64) ([]string, error) {
	var roleKeys []string

	// 构建查询
	err := dao.db.WithContext(ctx).
		// 从 sys_role 表开始，因为我们最终要取 role_key
		Table("sys_role r").
		// 连接 sys_user_role 表，通过 role_id 关联
		Joins("JOIN sys_user_role ur ON r.role_id = ur.role_id").
		// 连接 sys_user 表，通过 user_id 关联，用于筛选
		Joins("JOIN sys_user u ON ur.user_id = u.user_id").
		// 筛选出指定 userID 的记录
		Where("u.user_id = ?", userID).
		// 去重，防止因数据异常导致重复 (虽然通常不会)
		Distinct("r.role_key").
		// 提取 role_key 字段的值到 roleKeys 切片中
		Pluck("r.role_key", &roleKeys).Error

	// 如果查询出错（如数据库连接问题），返回错误
	if err != nil {
		return nil, fmt.Errorf("failed to query role keys for user %d: %w", userID, err)
	}

	// 查询成功，返回 roleKeys 列表
	// 注意：如果用户没有任何角色，roleKeys 将是空切片 []string{}，而不是 nil
	return roleKeys, nil
}

func (dao *SysMenuDAO) GetMenuTree(ctx context.Context) ([]*MenuTreeNode, error) {
	var objList []SysMenu
	err := dao.db.Find(&objList).Error
	if err != nil {
		return []*MenuTreeNode{}, err
	}

	tree := BuildMenuTree(objList)
	return tree, nil
}

func BuildMenuTree(objList []SysMenu) []*MenuTreeNode {
	// 1. 构建 map，用指针类型
	objMap := make(map[int64]*MenuTreeNode)
	for _, obj := range objList {
		objMap[obj.MenuID] = &MenuTreeNode{
			ID:       obj.MenuID,
			Label:    obj.MenuName,
			Disabled: false,
		}
	}

	// 2. 构建父子关系
	var roots []*MenuTreeNode
	for _, obj := range objList {
		node := objMap[obj.MenuID]

		if obj.ParentID == 0 {
			roots = append(roots, node)
		} else {
			if parent, exists := objMap[obj.ParentID]; exists {
				parent.Children = append(parent.Children, node)
			}
		}
	}

	// 3. 按 order_num 排序（递归排序）
	var sortNodes func([]*MenuTreeNode) []*MenuTreeNode
	sortNodes = func(nodes []*MenuTreeNode) []*MenuTreeNode {
		sort.Slice(nodes, func(i, j int) bool {
			return getMenuOrder(objList, nodes[i].ID) < getMenuOrder(objList, nodes[j].ID)
		})
		for i := range nodes {
			nodes[i].Children = sortNodes(nodes[i].Children)
		}
		return nodes
	}

	return sortNodes(roots)
}

func getMenuOrder(objList []SysMenu, id int64) int {
	for _, obj := range objList {
		if obj.MenuID == id {
			return obj.OrderNum
		}
	}
	return 9999
}
