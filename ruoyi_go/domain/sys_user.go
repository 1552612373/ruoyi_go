package domain

import "time"

type SysUser struct {
	// 用户ID
	// UserID int64 `gorm:"column:user_id;primaryKey;autoIncrement" json:"userId"`
	ID int64 `gorm:"column:user_id;primaryKey;autoIncrement" json:"userId"`

	// 部门ID
	DeptID *int64 `gorm:"column:dept_id" json:"deptId"`

	// 用户账号
	UserName string `gorm:"column:user_name" json:"userName"`

	// 用户昵称
	NickName string `gorm:"column:nick_name" json:"nickName"`

	// 用户类型（00系统用户）
	UserType string `gorm:"column:user_type" json:"userType"`

	// 用户邮箱
	Email string `gorm:"column:email" json:"email"`

	// 手机号码
	Phonenumber string `gorm:"column:phonenumber" json:"phonenumber"`

	// 用户性别（0男 1女 2未知）
	Sex string `gorm:"column:sex" json:"sex"`

	// 头像地址
	Avatar string `gorm:"column:avatar" json:"avatar"`

	// 密码
	Password string `gorm:"column:password" json:"password"`

	// 账号状态（0正常 1停用）
	Status string `gorm:"column:status" json:"status"`

	// 删除标志（0代表存在 2代表删除）
	DelFlag string `gorm:"column:del_flag" json:"delFlag"`

	// 最后登录IP
	LoginIP string `gorm:"column:login_ip" json:"loginIp"`

	// 最后登录时间（时间戳）
	LoginDate *time.Time `gorm:"column:login_date" json:"loginDate"`

	// 密码最后更新时间（时间戳）
	PwdUpdateDate *time.Time `gorm:"column:pwd_update_date" json:"pwdUpdateDate"`

	// 创建者
	CreateBy string `gorm:"column:create_by" json:"createBy"`

	// 创建时间（时间戳）
	CreateTime *time.Time `gorm:"column:create_time" json:"createTime"`

	// 更新者
	UpdateBy string `gorm:"column:update_by" json:"updateBy"`

	// 更新时间（时间戳）
	UpdateTime *time.Time `gorm:"column:update_time" json:"updateTime"`

	// 备注
	Remark *string `gorm:"column:remark" json:"remark"`

	Dept SysDept `json:"dept"`
}

type Params struct {
	BeginTime string `json:"beginTime" form:"params[beginTime]"` // URL 解码后是 params[beginTime]
	EndTime   string `json:"endTime" form:"params[endTime]"`
	// 或者如果希望直接解析为 time.Time 类型 (推荐用于日期比较)
	// BeginTime time.Time `json:"beginTime" form:"params[beginTime]"`
	// EndTime   time.Time `json:"endTime" form:"params[endTime]"`
}

// UserListReq 用于捕获完整的请求参数
type UserListReq struct {
	PageNum     int    `json:"pageNum" form:"pageNum"`
	PageSize    int    `json:"pageSize" form:"pageSize"`
	UserName    string `json:"userName" form:"userName"`       // 捕获 userName=aaa
	Phonenumber string `json:"phonenumber" form:"phonenumber"` // 捕获 phonenumber=13112341234
	Status      string `json:"status" form:"status"`           // 捕获 status=1
	Params      Params `json:"params" form:""`                 // 关键：这里 form 标签留空，Gin 会自动根据子字段的 form 标签去查找
}
