package acl

import (
	"gorm.io/gorm"
)

func RoleList(db *gorm.DB) []Role {
	var roleList []Role

	db.Find(&roleList)

	return roleList
}