package users

import(
	// "log"
	"github.com/Owicca/chan/models/acl"
	"gorm.io/gorm"
)

type User struct {
	ID int `gorm:"primaryKey;column:user_id"`
	DeletedAt int
	Username string
	Email string
	Password string
	Status string
	RoleId int
	Role acl.Role `gorm:"foreignKey:role_id;joinForeignKey:role_id"`
}

func GetUserList(db *gorm.DB) []User {
	var userList = []User{}

	db.Preload("Role").Find(&userList)

	return userList
}

func GetUser(db *gorm.DB, id int) User {
	var user User

	db.Preload("Role").First(&user, id)

	return user
}