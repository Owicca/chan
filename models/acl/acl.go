package acl

import (
	"gorm.io/gorm"
	// "log"
	"fmt"
)

var (
	ObjectActionList = make(ObjectAction)
	RoleActionObjectList = make(RolePair)
)

type ObjectAction map[int]string
type RolePair map[string][]string

type Action struct {
	ID int `gorm:"primaryKey;column:action_id"`
	DeletedAt int `gorm:"default=0"`
	Name string
	// Objects []*Object `gorm:"many2many:action_to_object;foreignKey:obj_id;joinForeignKey:obj_id;references:code;joinReferenceKey:code"`
}
type Object struct {
	ID int `gorm:"primaryKey;column:obj_id"`
	Name string
	// Actions []Action `gorm:"many2many:action_to_object;foreignKey:obj_id;joinForeignKey:obj_id;references:action_id;joinReferenceKey:action_id"`
}
type Role struct {
	ID int `gorm:"primaryKey;column:role_id"`
	DeletedAt int
	Name string
}

type ActionToObject struct{
	ID int `gorm:"primaryKey;column:ato_id"`
	ActionId int `gorm:"column:action_id"`
	ObjId int `gorm:"column:obj_id"`
	Action Action `gorm:"foreignKey:action_id;references:action_id"`
	Object Object `gorm:"foreignKey:obj_id;joinForeignKey:obj_id"`
}
func (ato ActionToObject) TableName() string {
	return "action_to_object"
}

func NewObjectAction(db *gorm.DB) ObjectAction {
	mp := make(ObjectAction)

	var atoList []ActionToObject
	db.Preload("Object").Preload("Action").Find(&atoList)

	for _, ato := range atoList {
		name := fmt.Sprintf("%s_%s", ato.Object.Name, ato.Action.Name)
		if _, ok := mp[ato.ID]; !ok {
			mp[ato.ID] = name
		}
	}

	return mp
}

type PairToRole struct{
	ID int `gorm:"primaryKey;column:ptr_id"`
	AtoId int `gorm:"column:ato_id"`
	RoleId int `gorm:"column:role_id"`
	ActionToObject ActionToObject `gorm:"foreignKey:ato_id;references:ato_id"`
	Role Role `gorm:"foreignKey:role_id;references:role_id"`
}
func (ptr PairToRole) TableName() string {
	return "pair_to_role"
}

// Only run this after a succesful run of NewObjectAction
func NewRolePair(db *gorm.DB) RolePair {
	mp := make(RolePair)

	var ptr []PairToRole
	db.Preload("Role").Preload("ActionToObject").Find(&ptr)

	for _, v := range ptr {
		if _, ok := mp[v.Role.Name]; !ok {
			mp[v.Role.Name] = []string{ObjectActionList[v.ActionToObject.ID]}
		} else {
			mp[v.Role.Name] = append(mp[v.Role.Name], ObjectActionList[v.ActionToObject.ID])
		}
	}

	return mp
}

func Run(db *gorm.DB) bool {
	result := false

	ObjectActionList = NewObjectAction(db)
	RoleActionObjectList = NewRolePair(db)

	return result
}