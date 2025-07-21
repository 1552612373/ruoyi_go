package dao

import "gorm.io/gorm"

func InitTable(db *gorm.DB) error {
	// 临时调试调用
	db.AutoMigrate(&SysDept{})
	db.AutoMigrate(&SysUser{})
	return db.AutoMigrate(&SysRole{})
}
