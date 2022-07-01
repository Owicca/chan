package acl

import (
	"gorm.io/gorm"
)

func RoleList(db *gorm.DB) []Role {
	var roleList []Role

	db.Not("name = ?", "root").Not("deleted_at > ?", 0).Find(&roleList)

	return roleList
}

func RoleIdList(db *gorm.DB) []int {
	var roleIdList []int

	db.Table("roles").Select("id").Not("name = ?", "root").Not("deleted_at > ?", 0).Scan(&roleIdList)

	return roleIdList
}
