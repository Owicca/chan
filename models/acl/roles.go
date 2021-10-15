package acl

import (
	"gorm.io/gorm"
	// "log"
)

func RoleList(db *gorm.DB) []Role {
	var roleList []Role

	db.Find(&roleList)

	return roleList
}