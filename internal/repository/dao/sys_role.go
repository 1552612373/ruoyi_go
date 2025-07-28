package dao

import (
	"context"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// SysRole 角色信息表(sys_role)
type SysRole struct {
	// 角色ID（主键）
	RoleId int64 `json:"roleId" gorm:"column:role_id;primaryKey;autoIncrement"`

	// 角色名称（如：管理员、普通用户）
	RoleName string `json:"roleName" gorm:"column:role_name"`

	// 角色权限字符串（如：admin、common）
	RoleKey string `json:"roleKey" gorm:"column:role_key"`

	// 角色排序（用于界面展示顺序）
	RoleSort int `json:"roleSort" gorm:"column:role_sort"`

	// 数据权限范围（如：1-全部数据，2-本部门，3-自定义）
	DataScope string `json:"dataScope" gorm:"column:data_scope"`

	// 菜单权限是否严格校验父子关系（0-否，1-是）
	MenuCheckStrictly int `json:"menuCheckStrictly" gorm:"column:menu_check_strictly"`

	// 部门权限是否严格校验父子关系（0-否，1-是）
	DeptCheckStrictly int `json:"deptCheckStrictly" gorm:"column:dept_check_strictly"`

	// 角色状态（0-启用，1-停用）
	Status string `json:"status" gorm:"column:status"`

	// 删除标志（0-正常，1-已删除）
	DelFlag string `json:"delFlag" gorm:"column:del_flag"`

	// 创建人（操作者用户名）
	CreateBy string `json:"createBy" gorm:"column:create_by"`

	// 创建时间
	CreateTime time.Time `json:"createTime" gorm:"column:create_time"`

	// 最后更新人
	UpdateBy string `json:"updateBy" gorm:"column:update_by"`

	// 最后更新时间
	UpdateTime time.Time `json:"updateTime" gorm:"column:update_time"`

	// 备注信息（可选）
	Remark string `json:"remark" gorm:"column:remark"`
}

type SysRoleMenu struct {
	RoleId int64 `gorm:"primaryKey"`
	MenuId int64 `gorm:"primaryKey"`
}

type SysRoleDept struct {
	RoleId int64 `gorm:"primaryKey"`
	DeptId int64 `gorm:"primaryKey"`
}

type SysRoleDAO struct {
	db      *gorm.DB
	menuDao *SysMenuDAO
}

func NewSysRoleDAO(db *gorm.DB, menuDao *SysMenuDAO) *SysRoleDAO {
	return &SysRoleDAO{
		db:      db,
		menuDao: menuDao,
	}
}

func (dao *SysRoleDAO) Insert(ctx context.Context, obj SysRole, menuIds []int64, deptIds []int64) error {
	// 开启事务
	tx := dao.db.WithContext(ctx).Begin()
	// “延迟执行 + panic 捕获” 机制，用于在发生 panic 时，自动回滚事务，防止数据不一致
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 先插入角色表
	err := dao.db.WithContext(ctx).Create(&obj).Error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflictsErrNo uint16 = 1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			// 唯一键冲突
			tx.Rollback()
			return errors.New("ZT唯一键冲突")
		}
	}
	// 如果有菜单，则插入关系表
	if len(menuIds) > 0 {
		var roleMenus []SysRoleMenu
		for _, menuId := range menuIds {
			roleMenus = append(roleMenus, SysRoleMenu{
				RoleId: obj.RoleId,
				MenuId: menuId,
			})
		}
		if err := tx.Create(&roleMenus).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	// 如果有部门，则插入关系表
	if len(deptIds) > 0 {
		var roleDepts []SysRoleDept
		for _, deptId := range deptIds {
			roleDepts = append(roleDepts, SysRoleDept{
				RoleId: obj.RoleId,
				DeptId: deptId,
			})
		}
		if err := tx.Create(&roleDepts).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return err
}

func (dao *SysRoleDAO) QueryList(ctx context.Context, pageNum int, pageSize int) ([]SysRole, int, error) {
	objList := []SysRole{}
	db := dao.db.WithContext(ctx).Model(&SysRole{})

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

func (dao *SysRoleDAO) Update(ctx context.Context, obj SysRole) error {
	err := dao.db.WithContext(ctx).Model(&obj).Where("role_id = ?", obj.RoleId).Updates(obj).Error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflictsErrNo uint16 = 1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			// 唯一键冲突
			return errors.New("ZT唯一键冲突")
		}
	}
	return err
}

func (dao *SysRoleDAO) QueryById(ctx context.Context, id int64) (SysRole, error) {
	obj := SysRole{}
	err := dao.db.WithContext(ctx).Where("role_id = ?", id).First(&obj)
	return obj, err.Error
}

func (dao *SysRoleDAO) DeleteById(ctx context.Context, id int64) error {
	// 开启事务
	tx := dao.db.WithContext(ctx).Begin()
	// “延迟执行 + panic 捕获” 机制，用于在发生 panic 时，自动回滚事务，防止数据不一致
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 先删角色
	err := dao.db.WithContext(ctx).Where("role_id = ?", id).Delete(&SysRole{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// 再删关系表
	if err := tx.Where("role_id = ?", id).Delete(&SysRoleMenu{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("role_id = ?", id).Delete(&SysRoleDept{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 提交事务
	return tx.Commit().Error
}

// 角色菜单树
func (dao *SysRoleDAO) QueryRoleMenuTreeById(ctx context.Context, id int64) ([]*MenuTreeNode, []int64, error) {
	// 所有菜单树形结构
	tree, err := dao.menuDao.GetMenuTree(ctx)
	if err != nil {
		return []*MenuTreeNode{}, []int64{}, err
	}

	// 角色菜单关系表
	var roleMenus []SysRoleMenu
	errx := dao.db.Where("role_id = ?", id).Find(&roleMenus).Error
	if errx != nil {
		return []*MenuTreeNode{}, []int64{}, err
	}
	var ids []int64
	for _, rel := range roleMenus {
		ids = append(ids, rel.MenuId)
	}

	return tree, ids, nil

}
