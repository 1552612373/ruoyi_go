package domain

type SysDictData struct {
	// 字典数据ID（主键）
	DictCode int64 `json:"dictCode" gorm:"column:dict_code;primaryKey"`

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

	ListClass string `json:"listClass" gorm:"column:list_class"`
	CssClass  string `json:"cssClass" gorm:"column:css_class"`
}
