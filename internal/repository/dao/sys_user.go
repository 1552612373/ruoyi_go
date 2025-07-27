package dao

import (
	"context"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type SysUser struct {
	// 用户ID
	UserID int64 `gorm:"column:user_id;primaryKey;autoIncrement" json:"userId"`

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
	LoginDate time.Time `gorm:"column:login_date" json:"loginDate"`

	// 密码最后更新时间（时间戳）
	PwdUpdateDate time.Time `gorm:"column:pwd_update_date" json:"pwdUpdateDate"`

	// 创建者
	CreateBy string `gorm:"column:create_by" json:"createBy"`

	// 创建时间（时间戳）
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`

	// 更新者
	UpdateBy string `gorm:"column:update_by" json:"updateBy"`

	// 更新时间（时间戳）
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`

	// 备注
	Remark *string `gorm:"column:remark" json:"remark"`
}

type SysUserDAO struct {
	db *gorm.DB
}

func NewSysUserDAO(db *gorm.DB) *SysUserDAO {
	return &SysUserDAO{
		db: db,
	}
}

func (dao *SysUserDAO) Insert(ctx context.Context, obj SysUser) error {
	now := time.Now()
	obj.UpdateTime = now
	obj.CreateTime = now
	err := dao.db.WithContext(ctx).Create(&obj).Error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflictsErrNo uint16 = 1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			// 唯一键冲突
			return errors.New("ZT账户已存在")
		}
	}
	return err
}

func (dao *SysUserDAO) FindByAccount(ctx context.Context, account string) (SysUser, error) {
	sysUser := SysUser{}
	err := dao.db.WithContext(ctx).Where("user_name = ?", account).First(&sysUser).Error
	return sysUser, err
}

func (dao *SysUserDAO) FindById(ctx context.Context, id int64) (SysUser, error) {
	sysUser := SysUser{}
	err := dao.db.WithContext(ctx).Where("user_id = ?", id).First(&sysUser).Error
	return sysUser, err
}

func (dao *SysUserDAO) QueryList(ctx context.Context, pageNum int, pageSize int) ([]SysUser, int, error) {
	objList := []SysUser{}
	db := dao.db.WithContext(ctx).Model(&SysUser{})

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
