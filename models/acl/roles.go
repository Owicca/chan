package acl

import (
	"gorm.io/gorm"
)

var (
	privilegedRoles = []string{"root", "board_admin", "op"}
)

func RoleList(db *gorm.DB) []Role {
	var roleList []Role

	db.Where("name NOT IN (?)", privilegedRoles).Not("deleted_at > ?", 0).Find(&roleList)

	return roleList
}

func RoleIdList(db *gorm.DB) []int {
	var roleIdList []int

	db.Table("roles").Select("id").Where("name NOT IN (?)", privilegedRoles).Not("deleted_at > ?", 0).Scan(&roleIdList)

	return roleIdList
}
