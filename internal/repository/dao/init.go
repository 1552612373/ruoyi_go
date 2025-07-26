package dao

import "gorm.io/gorm"

func InitTable(db *gorm.DB) error {
	// 临时调试调用
	db.AutoMigrate(&SysDept{})
	db.AutoMigrate(&SysUser{})
	db.AutoMigrate(&SysDictType{})
	db.AutoMigrate(&SysDictData{})
	db.AutoMigrate(&SysUser{})
	db.AutoMigrate(&SysPost{})
	db.AutoMigrate(&SysMenu{})
	return db.AutoMigrate(&SysRole{})
}
