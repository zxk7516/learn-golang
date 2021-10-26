package models

import "zxk/go-fiber-demo/common"

func MakeMigrations() {
	common.DB.AutoMigrate(&User{})
	common.DB.AutoMigrate(&AuthToken{})
	common.DB.AutoMigrate(&Page{})

	common.DB.AutoMigrate(&AdminUsers{})
	common.DB.AutoMigrate(&AdminRole{})
	common.DB.AutoMigrate(&AdminPersmisions{})
	common.DB.AutoMigrate(&AdminMenu{})
	common.DB.AutoMigrate(&AdminRoleUser{})
	common.DB.AutoMigrate(&AdminRolePermisions{})
	common.DB.AutoMigrate(&AdminUserPermision{})
	common.DB.AutoMigrate(&AdminRoleMenu{})

	common.DB.AutoMigrate(&AdminOperationLog{})

	InitAdminTables()

}

func InitAdminTables() {
	common.DB.Exec("truncate TABLE admin_roles;")
	adminRole := AdminRole{Name: "admin", Slug: "admin"}
	common.DB.Create(&adminRole)

}
