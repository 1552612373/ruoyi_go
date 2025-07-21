package domain

type SysRole struct {
	// 角色ID（主键）
	RoleId int64 `json:"roleId" gorm:"column:role_id;primaryKey"`

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
	CreateTime int64 `json:"createTime" gorm:"column:create_time"`

	// 最后更新人
	UpdateBy string `json:"updateBy" gorm:"column:update_by"`

	// 最后更新时间
	UpdateTime int64 `json:"updateTime" gorm:"column:update_time"`

	// 备注信息（可选）
	Remark string `json:"remark" gorm:"column:remark"`
}
