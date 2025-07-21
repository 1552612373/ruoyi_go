package domain

type SysUser struct {
	// 用户ID（主键）
	UserId int64 `json:"userId" gorm:"column:user_id;primaryKey"`

	// 所属部门ID
	DeptId int64 `json:"deptId" gorm:"column:dept_id"`

	// 登录账号（唯一）
	UserName string `json:"userName" gorm:"column:user_name"`

	// 用户昵称（显示用）
	NickName string `json:"nickName" gorm:"column:nick_name"`

	// 邮箱地址
	Email string `json:"email" gorm:"column:email"`

	// 头像地址（URL 或 Base64）
	Avatar string `json:"avatar" gorm:"column:avatar"`

	// 手机号码
	Phonenumber string `json:"phonenumber" gorm:"column:phonenumber"`

	// 性别 男、女
	Sex string `json:"sex" gorm:"column:sex"`

	// 登录密码（加密存储）
	Password string `json:"password" gorm:"column:password"`

	// 账户状态（0-启用，1-停用）
	Status string `json:"status" gorm:"column:status"`

	// 删除标志（0-正常，1-已删除，2-彻底删除）
	DelFlag string `json:"delFlag" gorm:"column:del_flag"`

	// 最后登录IP
	LoginIp string `json:"loginIp" gorm:"column:login_ip"`

	// 最后登录时间
	LoginDate int64 `json:"loginDate" gorm:"column:login_date"`

	// 密码最后修改时间
	PwdUpdateDate int64 `json:"pwdUpdateDate" gorm:"column:pwd_update_date"`

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

	// 关联部门信息（一对一）
	Dept *SysDept `json:"dept,omitempty" gorm:"foreignKey:DeptId"`

	// 用户拥有的角色列表（多对多）
	Roles []SysRole `json:"roles,omitempty" gorm:"many2many:sys_user_role;"`
}
