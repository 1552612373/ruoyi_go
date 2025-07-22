package dao

import "gorm.io/gorm"

func InitTable(db *gorm.DB) error {
	// 临时调试调用
	db.AutoMigrate(&SysDept{})
	db.AutoMigrate(&SysUser{})
	err := db.AutoMigrate(&SysDictType{})
	if err != nil {
		print(err)
	}
	db.AutoMigrate(&SysDictData{})
	db.AutoMigrate(&SysUser{})
	return db.AutoMigrate(&SysRole{})
}
