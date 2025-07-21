package dao

import "time"

// SysDept 部门信息表(sys_dept)
type SysDept struct {
	// 部门ID（主键）
	DeptId int64 `json:"deptId" gorm:"column:dept_id;primaryKey"`

	// 父部门ID（顶级部门为0）
	ParentId int64 `json:"parentId" gorm:"column:parent_id"`

	// 祖级部门ID集合（逗号分隔，如：0,1,2）
	Ancestors string `json:"ancestors" gorm:"column:ancestors"`

	// 部门名称
	DeptName string `json:"deptName" gorm:"column:dept_name"`

	// 显示顺序（用于排序）
	OrderNum int `json:"orderNum" gorm:"column:order_num"`

	// 负责人名称
	Leader string `json:"leader" gorm:"column:leader"`

	// 联系电话
	Phone string `json:"phone" gorm:"column:phone"`

	// 邮箱
	Email string `json:"email" gorm:"column:email"`

	// 部门状态（0-正常，1-停用）
	Status string `json:"status" gorm:"column:status"`

	// 删除标志（0-正常，1-已删除）
	DelFlag string `json:"delFlag" gorm:"column:del_flag"`

	// 父部门名称（非数据库字段，用于展示）
	ParentName string `json:"parentName" gorm:"-"`

	// 创建人（操作者用户名）
	CreateBy string `json:"createBy" gorm:"column:create_by"`

	// 创建时间
	CreateTime time.Time `json:"createTime" gorm:"column:create_time"`

	// 最后更新人
	UpdateBy string `json:"updateBy" gorm:"column:update_by"`

	// 最后更新时间
	UpdateTime time.Time `json:"updateTime" gorm:"column:update_time"`
}
