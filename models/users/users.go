package users

import(
	"github.com/Owicca/chan/models/acl"
	"gorm.io/gorm"
)

type User struct {
	ID int `gorm:"primaryKey;column:id"`
	Deleted_at int64
	Username string
	Email string
	Password string
	Status string
	RoleId int
	Role acl.Role `gorm:"foreignKey:role_id;"`
}

func UserList(db *gorm.DB) []User {
	userList := []User{}

	db.Preload("Role").Find(&userList)

	return userList
}

func UserOne(db *gorm.DB, id int) User {
	user := User{}

	db.Preload("Role").First(&user, id)

	return user
}

func UserOneCreate(db *gorm.DB, user User) error {
	return db.Create(&user).Error
}

func UserOneUpdate(db *gorm.DB, user User) error {
	return db.Model(&user).Updates(user).Error
}