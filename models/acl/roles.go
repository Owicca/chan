package acl

import (
	"gorm.io/gorm"
	// "log"
)

func GetRoleList(db *gorm.DB) []Role {
	var roleList []Role

	db.Find(&roleList)

	return roleList
}