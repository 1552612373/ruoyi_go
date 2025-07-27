package dao

import (
	"context"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type SysPost struct {
	// 岗位ID
	PostID int64 `gorm:"column:post_id;primaryKey;autoIncrement" json:"postId"`
	// 岗位编码
	PostCode string `gorm:"column:post_code" json:"postCode"`
	// 岗位名称
	PostName string `gorm:"column:post_name" json:"postName"`
	// 显示顺序
	PostSort int32 `gorm:"column:post_sort" json:"postSort"`
	// 状态（0正常 1停用）
	Status string `gorm:"column:status" json:"status"`
	// 创建者
	CreateBy string `gorm:"column:create_by" json:"createBy"`
	// 创建时间
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	// 更新者
	UpdateBy string `gorm:"column:update_by" json:"updateBy"`
	// 更新时间
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	// 备注
	Remark *string `gorm:"column:remark" json:"remark"`
}

type SysPostDAO struct {
	db *gorm.DB
}

func NewSysPostDAO(db *gorm.DB) *SysPostDAO {
	return &SysPostDAO{
		db: db,
	}
}

func (dao *SysPostDAO) Insert(ctx context.Context, obj SysPost) error {
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

func (dao *SysPostDAO) QueryList(ctx context.Context, pageNum int, pageSize int) ([]SysPost, int, error) {
	objList := []SysPost{}
	db := dao.db.WithContext(ctx).Model(&SysPost{})

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

func (dao *SysPostDAO) Update(ctx context.Context, obj SysPost) error {
	err := dao.db.WithContext(ctx).Model(&obj).Where("post_id = ?", obj.PostID).Updates(obj).Error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflictsErrNo uint16 = 1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			// 唯一键冲突
			return errors.New("ZT唯一键冲突")
		}
	}
	return err
}

func (dao *SysPostDAO) QueryById(ctx context.Context, id int64) (SysPost, error) {
	obj := SysPost{}
	err := dao.db.WithContext(ctx).Where("post_id = ?", id).First(&obj)
	return obj, err.Error
}

func (dao *SysPostDAO) DeleteById(ctx context.Context, id int64) error {
	err := dao.db.WithContext(ctx).Where("post_id = ?", id).Delete(&SysPost{}).Error
	return err
}
