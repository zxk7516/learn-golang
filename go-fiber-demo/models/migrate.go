package models

import (
	"zxk/go-fiber-demo/common"
	"zxk/go-fiber-demo/models/admin"
)

func MakeMigrations() {
	common.DB.AutoMigrate(&User{})
	common.DB.AutoMigrate(&AuthToken{})
	common.DB.AutoMigrate(&Page{})

	common.DB.AutoMigrate(&admin.AdminUsers{})
	common.DB.AutoMigrate(&admin.AdminRole{})
	common.DB.AutoMigrate(&admin.AdminPersmisions{})
	common.DB.AutoMigrate(&admin.AdminMenu{})
	common.DB.AutoMigrate(&admin.AdminRoleUser{})
	common.DB.AutoMigrate(&admin.AdminRolePermisions{})
	common.DB.AutoMigrate(&admin.AdminUserPermision{})
	common.DB.AutoMigrate(&admin.AdminRoleMenu{})

	common.DB.AutoMigrate(&admin.AdminOperationLog{})

	InitAdminTables()

}

func InitAdminTables() {
	common.DB.Exec("truncate TABLE admin_roles;")
	adminRole := admin.AdminRole{Name: "admin", Slug: "admin"}
	common.DB.Create(&adminRole)
	common.DB.Exec("truncate TABLE admin_users;")
	adminUser := admin.AdminUsers{UserName: "admin", Password: common.HashPassword("admin"), Name: "管理员"}
	common.DB.Create(&adminUser)
	common.DB.Exec("truncate TABLE admin_role_users;")
	adminRoleUser := admin.AdminRoleUser{UserId: adminUser.ID, RoleId: adminRole.ID}
	common.DB.Create(adminRoleUser)
}
