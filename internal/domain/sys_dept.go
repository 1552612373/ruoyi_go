package domain

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
	CreateTime int64 `gorm:"column:create_time" json:"createTime"`

	// 更新者
	UpdateBy string `gorm:"column:update_by" json:"updateBy"`

	// 更新时间（时间戳）
	UpdateTime int64 `gorm:"column:update_time" json:"updateTime"`
}
