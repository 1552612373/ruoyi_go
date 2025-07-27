package dao

import (
	"context"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type SysDictData struct {
	// 字典编码
	DictCode int64 `gorm:"column:dict_code;primaryKey;autoIncrement" json:"dictCode"`

	// 字典排序
	DictSort int `gorm:"column:dict_sort" json:"dictSort"`

	// 字典标签
	DictLabel string `gorm:"column:dict_label" json:"dictLabel"`

	// 字典键值
	DictValue string `gorm:"column:dict_value" json:"dictValue"`

	// 字典类型
	DictType string `gorm:"column:dict_type" json:"dictType"`

	// 样式属性（其他样式扩展）
	CssClass *string `gorm:"column:css_class" json:"cssClass"`

	// 表格回显样式
	ListClass *string `gorm:"column:list_class" json:"listClass"`

	// 是否默认（Y是 N否）
	IsDefault string `gorm:"column:is_default" json:"isDefault"`

	// 状态（0正常 1停用）
	Status string `gorm:"column:status" json:"status"`

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

type SysDictDataDAO struct {
	db *gorm.DB
}

func NewSysDictDataDAO(db *gorm.DB) *SysDictDataDAO {
	return &SysDictDataDAO{
		db: db,
	}
}

func (dao *SysDictDataDAO) Insert(ctx context.Context, obj SysDictData) error {
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

func (dao *SysDictDataDAO) QueryList(ctx context.Context, pageNum int, pageSize int, dictType string) ([]SysDictData, int, error) {
	objList := []SysDictData{}
	db := dao.db.WithContext(ctx).Model(&SysDictData{})

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
	err := db.Where("dict_type = ?", dictType).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&objList).Error

	return objList, int(total), err
}

func (dao *SysDictDataDAO) QueryByDictCode(ctx context.Context, dictCode int64) (SysDictData, error) {
	obj := SysDictData{}
	err := dao.db.WithContext(ctx).Where("dict_code = ?", dictCode).First(&obj)
	return obj, err.Error
}

func (dao *SysDictDataDAO) Update(ctx context.Context, obj SysDictData) error {
	err := dao.db.WithContext(ctx).Model(&obj).Where("dict_code = ?", obj.DictCode).Updates(obj).Error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflictsErrNo uint16 = 1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			// 唯一键冲突
			return errors.New("ZT唯一键冲突")
		}
	}
	return err
}

func (dao *SysDictDataDAO) DeleteByDictCode(ctx context.Context, dictCode int64) error {
	err := dao.db.WithContext(ctx).Where("dict_code = ?", dictCode).Delete(&SysDictData{}).Error
	return err
}
