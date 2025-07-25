package domain

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
	CreateTime int64 `gorm:"column:create_time" json:"createTime"`
	// 更新者
	UpdateBy string `gorm:"column:update_by" json:"updateBy"`
	// 更新时间
	UpdateTime int64 `gorm:"column:update_time" json:"updateTime"`
	// 备注
	Remark *string `gorm:"column:remark" json:"remark"`
}
