package users

import(
	// "log"
	"github.com/Owicca/chan/models/acl"
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