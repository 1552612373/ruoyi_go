package domain

import "time"

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
	CreateTime time.Time `json:"createTime" gorm:"column:create_time"`

	// 更新者
	UpdateBy string `json:"updateBy" gorm:"column:update_by"`

	// 更新时间
	UpdateTime time.Time `json:"updateTime" gorm:"column:update_time"`

	// 备注
	Remark string `json:"remark" gorm:"column:remark"`

	// 关联的字典数据（可选，用于关联查询）
	// DictDataList []SysDictData `json:"dictDataList,omitempty" gorm:"foreignKey:DictType;references:DictType"`
}
