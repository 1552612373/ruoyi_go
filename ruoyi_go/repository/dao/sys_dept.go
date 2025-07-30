package dao

import (
	"context"
	"errors"
	"sort"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type SysDept struct {
	// 部门id
	DeptID int64 `gorm:"column:dept_id;primaryKey;autoIncrement" json:"deptId"`

	// 父部门id
	ParentID int64 `gorm:"column:parent_id" json:"parentId"`

	// 祖级列表
	Ancestors string `gorm:"column:ancestors" json:"ancestors"`

	// 部门名称
	DeptName string `gorm:"column:dept_name" json:"deptName"`

	// 显示顺序
	OrderNum int `gorm:"column:order_num" json:"orderNum"`

	// 负责人
	Leader *string `gorm:"column:leader" json:"leader"`

	// 联系电话
	Phone *string `gorm:"column:phone" json:"phone"`

	// 邮箱
	Email *string `gorm:"column:email" json:"email"`

	// 部门状态（0正常 1停用）
	Status string `gorm:"column:status" json:"status"`

	// 删除标志（0代表存在 2代表删除）
	DelFlag string `gorm:"column:del_flag" json:"delFlag"`

	// 创建者
	CreateBy string `gorm:"column:create_by" json:"createBy"`

	// 创建时间（时间戳）
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`

	// 更新者
	UpdateBy string `gorm:"column:update_by" json:"updateBy"`

	// 更新时间（时间戳）
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
}

type DeptTreeNode struct {
	ID       int64           `json:"id"`
	Label    string          `json:"label"`
	Disabled bool            `json:"disabled"`
	Children []*DeptTreeNode `json:"children,omitempty"`
}

type SysDeptDAO struct {
	db *gorm.DB
}

func NewSysDeptDAO(db *gorm.DB) *SysDeptDAO {
	return &SysDeptDAO{
		db: db,
	}
}

func (dao *SysDeptDAO) Insert(ctx context.Context, obj SysDept) error {
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

func (dao *SysDeptDAO) QueryList(ctx context.Context) ([]SysDept, int, error) {
	objList := []SysDept{}
	db := dao.db.WithContext(ctx).Model(&SysDept{})

	var total int64
	var pageNum = 1
	var pageSize = 1000

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

func (dao *SysDeptDAO) QueryListExclude(ctx context.Context, excludeDeptId int64) ([]SysDept, int, error) {
	objList := []SysDept{}
	db := dao.db.WithContext(ctx).Model(&SysDept{})

	var total int64
	var pageNum = 1
	var pageSize = 1000

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
	err := db.Where("dept_id <> ?", excludeDeptId).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&objList).Error

	return objList, int(total), err
}

func (dao *SysDeptDAO) QueryByDeptId(ctx context.Context, deptId int64) (SysDept, error) {
	obj := SysDept{}
	err := dao.db.WithContext(ctx).Where("dept_id = ?", deptId).First(&obj)
	return obj, err.Error
}

func (dao *SysDeptDAO) DeleteById(ctx context.Context, id int64) error {
	err := dao.db.WithContext(ctx).Where("dept_id = ?", id).Delete(&SysDept{}).Error
	return err
}

func (dao *SysDeptDAO) Update(ctx context.Context, obj SysDept) error {
	err := dao.db.WithContext(ctx).Model(&obj).Where("dept_id = ?", obj.DeptID).Updates(obj).Error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflictsErrNo uint16 = 1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			// 唯一键冲突
			return errors.New("ZT唯一键冲突")
		}
	}
	return err
}

func (dao *SysDeptDAO) GetDeptTree(ctx context.Context) ([]*DeptTreeNode, error) {
	var depts []SysDept
	err := dao.db.Find(&depts).Error
	if err != nil {
		return []*DeptTreeNode{}, err
	}

	tree := BuildDeptTree(depts)
	return tree, nil
}

func BuildDeptTree(depts []SysDept) []*DeptTreeNode {
	// 1. 构建 map，用指针类型
	deptMap := make(map[int64]*DeptTreeNode)
	for _, dept := range depts {
		deptMap[dept.DeptID] = &DeptTreeNode{
			ID:       dept.DeptID,
			Label:    dept.DeptName,
			Disabled: false,
		}
	}

	// 2. 构建父子关系
	var roots []*DeptTreeNode
	for _, dept := range depts {
		node := deptMap[dept.DeptID]

		if dept.ParentID == 0 {
			roots = append(roots, node)
		} else {
			if parent, exists := deptMap[dept.ParentID]; exists {
				parent.Children = append(parent.Children, node)
			}
		}
	}

	// 3. 按 order_num 排序（递归排序）
	var sortNodes func([]*DeptTreeNode) []*DeptTreeNode
	sortNodes = func(nodes []*DeptTreeNode) []*DeptTreeNode {
		sort.Slice(nodes, func(i, j int) bool {
			return getDeptOrder(depts, nodes[i].ID) < getDeptOrder(depts, nodes[j].ID)
		})
		for i := range nodes {
			nodes[i].Children = sortNodes(nodes[i].Children)
		}
		return nodes
	}

	return sortNodes(roots)
}

func getDeptOrder(depts []SysDept, id int64) int {
	for _, dept := range depts {
		if dept.DeptID == id {
			return dept.OrderNum
		}
	}
	return 9999
}
