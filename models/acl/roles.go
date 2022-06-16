package acl

import (
	"gorm.io/gorm"
)

func RoleList(db *gorm.DB) []Role {
	var roleList []Role

	db.Find(&roleList)

	return roleList
}

func RoleIdList(db *gorm.DB) []int {
	var roleIdList []int

	db.Raw(`
	SELECT id FROM roles;
	`).Scan(&roleIdList)

	return roleIdList
}
