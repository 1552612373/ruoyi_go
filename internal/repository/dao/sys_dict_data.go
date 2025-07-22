package dao

import (
	"context"
	"errors"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// SysDictData 字典数据表(sys_dict_data)
type SysDictData struct {
	// 字典数据ID（主键）
	DictCode int64 `json:"dictCode" gorm:"column:dict_code;primaryKey;autoIncrement"`

	// 字典排序
	DictSort int `json:"dictSort" gorm:"column:dict_sort"`

	// 字典标签（如：男、女、保密）
	DictLabel string `json:"dictLabel" gorm:"column:dict_label"`

	// 字典值（如：0、1、2）
	DictValue string `json:"dictValue" gorm:"column:dict_value"`

	// 字典类型编码（如：sys_user_sex）
	DictType string `json:"dictType" gorm:"column:dict_type;index"`

	// 状态（0正常 1停用）
	Status string `json:"status" gorm:"column:status"`

	// 创建者
	CreateBy string `json:"createBy" gorm:"column:create_by"`

	// 创建时间
	CreateTime int64 `json:"createTime" gorm:"column:create_time"`

	// 更新者
	UpdateBy string `json:"updateBy" gorm:"column:update_by"`

	// 更新时间
	UpdateTime int64 `json:"updateTime" gorm:"column:update_time"`

	// 备注
	Remark string `json:"remark" gorm:"column:remark"`
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
