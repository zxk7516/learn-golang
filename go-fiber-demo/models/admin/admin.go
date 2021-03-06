package admin

import (
	"zxk/go-fiber-demo/common"
)

/*
AdminUsers
drop table admin_menus ;
drop table admin_persmisions ;
drop table admin_role_permisions
drop table admin_role_users
drop table admin_users
drop table admin_operation_logs
drop table admin_role_menus
drop table admin_roles
drop table admin_user_permisions
*/
type AdminUsers struct {
	ID       uint64 `gorm:"column:id;primary_key" json:"id"`
	UserName string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"-"`
	Name     string `gorm:"column:name" json:"name"`
	Avator   string `gorm:"column:avator" json:"avator"`
	common.TimeStamps
}

type AdminRole struct {
	ID   uint64 `gorm:"column:id;primary_key" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Slug string `gorm:"column:slug" json:"slug"`
	common.TimeStamps
}

type AdminPersmisions struct {
	ID         uint64 `gorm:"column:id;primary_key" json:"id"`
	Name       string `gorm:"column:name" json:"name"`
	Slug       string `gorm:"column:slug" json:"slug"`
	HttpMethod string `gorm:"column:http_method" json:"http_method"`
	HttpPath   string `gorm:"column:http_path" json:"http_path"`
	common.TimeStamps
}

type AdminMenu struct {
	ID         uint64 `gorm:"column:id;primary_key" json:"id"`
	ParentId   uint64 `gorm:"column:parent_id" json:"parent_id"`
	Order      int    `gorm:"column:order" json:"order"`
	Title      string `gorm:"column:title" json:"title"`
	Icon       string `gorm:"column:icon" json:"icon"`
	Uri        string `gorm:"column:uri" json:"uri"`
	Permisions string `gorm:"column:permisions" json:"permisions"`
	common.TimeStamps
}

type AdminRoleUser struct {
	UserId uint64 `gorm:"column:user_id;index:idx_user_role" json:"user_id"`
	RoleId uint64 `gorm:"column:role_id;index:idx_user_role" json:"role_id"`
}

type AdminRolePermisions struct {
	RoleId      uint64 `gorm:"column:role_id;index:idx_role_permision" json:"role_id"`
	PermisionId uint64 `gorm:"column:permision_id;index:idx_role_permision" json:"permision_id"`
}

type AdminUserPermision struct {
	UserId      uint64 `gorm:"column:user_id;index:idx_user_permision" json:"user_id"`
	PermisionId uint64 `gorm:"column:permision_id;index:idx_user_permision" json:"permision_id"`
}

type AdminRoleMenu struct {
	RoleId uint64 `gorm:"column:role_id;index:idx_role_menu" json:"role_id"`
	MenuId uint64 `gorm:"column:menu_id;index:idx_role_menu" json:"menu_id"`
}

type AdminOperationLog struct {
	ID         uint64 `gorm:"column:id;primary_key" json:"id"`
	UserId     uint64 `gorm:"column:user_id;index" json:"user_id"`
	HttpMethod string `gorm:"column:http_method" json:"http_method"`
	HttpPath   string `gorm:"column:http_path" json:"http_path"`
	Ip         string `gorm:"column:ip" json:"ip"`
	Input      string `gorm:"column:input" json:"input"`
	common.TimeStamps
}

func GetAdminUser(condition interface{}) (AdminUsers, error) {
	var user AdminUsers
	err := common.DB.Where(condition).First(&user).Error
	return user, err
}
