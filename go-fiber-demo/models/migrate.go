package models

import "zxk/go-fiber-demo/common"


func MakeMigrations() {
	common.DB.AutoMigrate(&User{})
	common.DB.AutoMigrate(&AuthToken{})
	common.DB.AutoMigrate(&Page{})
}
