package acl

import (
	"gorm.io/gorm"
	"fmt"
)

var (
	ObjectActionList = make(ObjectAction)
	RoleActionObjectList = make(RolePair)
)

type ObjectAction map[int]string
type RolePair map[string][]string

type Action struct {
	ID int `gorm:"primaryKey;column:id"`
	DeletedAt int64
	Name string
}
type Object struct {
	ID int `gorm:"primaryKey;column:id"`
	Name string
}
type Role struct {
	ID int `gorm:"primaryKey;column:id"`
	DeletedAt int64
	Name string
}

type ActionToObject struct{
	ID int `gorm:"primaryKey;column:id"`
	ActionId int `gorm:"column:action_id"`
	ObjId int `gorm:"column:obj_id"`
	Action Action `gorm:"foreignKey:action_id;references:id"`
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
	ID int `gorm:"primaryKey;column:id"`
	AtoId int `gorm:"column:ato_id"`
	RoleId int `gorm:"column:role_id"`
	ActionToObject ActionToObject `gorm:"foreignKey:ato_id;references:id"`
	Role Role `gorm:"foreignKey:role_id;references:id"`
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
	if len(ObjectActionList) > 0 || len(RoleActionObjectList) > 0 {
		result = true
	}

	return result
}

type Acl struct {
}

func Check(subject string, action string, object string, checkArr RolePair) bool {
	result := false

	if subject == "" || action == "" || object == "" {
		return result
	}

	objectAction := fmt.Sprintf("%s_%s", object, action)
	if arr, ok := checkArr[subject]; ok {
		for _, v := range arr {
			if v == objectAction {
				result = true
				break
			}
		}
	}

	return result
}