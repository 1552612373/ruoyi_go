package domain

import "time"

// SysMenu 菜单权限表
type SysMenu struct {
	// 菜单ID
	MenuID int64 `json:"menuId" gorm:"column:menu_id;primaryKey;autoIncrement"`

	// 菜单名称
	MenuName string `json:"menuName" gorm:"column:menu_name"`

	// 父菜单ID
	ParentID int64 `json:"parentId" gorm:"column:parent_id;default:0"`

	// 显示顺序
	OrderNum int `json:"orderNum" gorm:"column:order_num;default:0"`

	// 路由地址
	Path string `json:"path" gorm:"column:path;default:''"`

	// 组件路径
	Component string `json:"component" gorm:"column:component;default:null"`

	// 路由参数
	Query string `json:"query" gorm:"column:query;default:null"`

	// 路由名称
	RouteName string `json:"routeName" gorm:"column:route_name;default:''"`

	// 是否为外链（0是 1否）
	IsFrame int `json:"isFrame" gorm:"column:is_frame;default:1"`

	// 是否缓存（0缓存 1不缓存）
	IsCache int `json:"isCache" gorm:"column:is_cache;default:0"`

	// 菜单类型（M目录 C菜单 F按钮）
	MenuType string `json:"menuType" gorm:"column:menu_type;default:''"`

	// 菜单状态（0显示 1隐藏）
	Visible string `json:"visible" gorm:"column:visible;default:0"`

	// 菜单状态（0正常 1停用）
	Status string `json:"status" gorm:"column:status;default:0"`

	// 权限标识
	Perms string `json:"perms" gorm:"column:perms;default:null"`

	// 菜单图标
	Icon string `json:"icon" gorm:"column:icon;default:'#'"`

	// 创建者
	CreateBy string `json:"createBy" gorm:"column:create_by;default:''"`

	// 创建时间
	CreateTime time.Time `json:"createTime" gorm:"column:create_time"`

	// 更新者
	UpdateBy string `json:"updateBy" gorm:"column:update_by;default:''"`

	// 更新时间
	UpdateTime time.Time `json:"updateTime" gorm:"column:update_time"`

	// 备注
	Remark string `json:"remark" gorm:"column:remark;default:''"`
}
