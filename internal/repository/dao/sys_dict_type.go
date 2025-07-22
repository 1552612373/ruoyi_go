package dao

import (
	"context"
	"errors"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// SysDictType 字典类型表(sys_dict_type)
type SysDictType struct {
	// 字典类型ID（主键）
	DictId int64 `json:"dictId" gorm:"column:dict_id;primaryKey"`

	// 字典名称（如：用户性别）
	DictName string `json:"dictName" gorm:"column:dict_name"`

	// 字典类型编码（如：sys_user_sex）
	DictType string `json:"dictType" gorm:"column:dict_type;uniqueIndex;type:varchar(255)"`

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

	// 关联的字典数据（可选，用于关联查询）
	// DictDataList []SysDictData `json:"dictDataList,omitempty" gorm:"foreignKey:DictType;references:DictType"`
}

type SysDictTypeDAO struct {
	db *gorm.DB
}

func NewSysDictTypeDAO(db *gorm.DB) *SysDictTypeDAO {
	return &SysDictTypeDAO{
		db: db,
	}
}

func (dao *SysDictTypeDAO) Insert(ctx context.Context, obj SysDictType) error {
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

// func (dao *SysDictTypeDAO) QueryList(ctx context.Context, pageNum int, pageSize int) ([]SysDictType,error,int) {
// 	objList := []SysDictType{}
// 	dao.db.WithContext(ctx).Find(&objList)
// 	return err
// }
