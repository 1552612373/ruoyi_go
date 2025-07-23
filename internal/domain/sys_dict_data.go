package domain

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
	CreateTime int64 `gorm:"column:create_time" json:"createTime"`

	// 更新者
	UpdateBy string `gorm:"column:update_by" json:"updateBy"`

	// 更新时间（时间戳）
	UpdateTime int64 `gorm:"column:update_time" json:"updateTime"`

	// 备注
	Remark *string `gorm:"column:remark" json:"remark"`
}
