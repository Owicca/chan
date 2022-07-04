package users

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"

	"github.com/Owicca/chan/infra"
	"github.com/Owicca/chan/models/acl"
	"gorm.io/gorm"
	"upspin.io/errors"
)

const (
	UserPassMin = 1
	UserPassMax = 50
)

type User struct {
	ID         int `gorm:"primaryKey;column:id"`
	Deleted_at int64
	Username   string
	Email      string
	Password   string
	Pepper     string
	Status     string
	RoleId     int
	Role       acl.Role `gorm:"foreignKey:role_id;"`
}

func UserList(db *gorm.DB, limit int, offset int) []User {
	var userList []User

	stmt := db.Preload("Role")
	if limit > 0 {
		stmt = stmt.Limit(limit)
	}
	if offset > 0 {
		stmt = stmt.Offset(offset)
	}

	stmt.Find(&userList)

	return userList
}

func UserListCount(db *gorm.DB) int {
	var count int

	db.Raw(`
	SELECT COUNT(id) FROM users
	WHERE deleted_at = 0;
	`).Scan(&count)

	return count
}

func TotalActiveUsers(db *gorm.DB) int {
	var count int

	db.Raw(`
	SELECT COUNT(u.id) FROM users u
	LEFT JOIN roles r ON r.id = u.role_id
	WHERE
	r.name = ?
	`, "anon").Scan(&count)

	return count
}

func UserOne(db *gorm.DB, id int) User {
	user := User{}

	db.Preload("Role").First(&user, id)

	return user
}

func UserOneByEmail(db *gorm.DB, email string) User {
	user := User{}

	db.Preload("Role").First(&user, "email = ?", email)

	return user
}

func UserValidate(email string, pass1 string, pass2 string) error {
	if email == "" || pass1 == "" || pass2 == "" {
		return errors.Str("Email and password are required!")
	}

	rule := `.*\@.*\..*`
	emailReg := regexp.MustCompile(rule)
	if !emailReg.MatchString(email) {
		return errors.Str("Provided email is not valid!")
	}
	if pass1 != pass2 {
		return errors.Str("Passwords do not match!")
	}
	ln := len(pass1)
	if ln < UserPassMin || ln > UserPassMax {
		return errors.Str("Invalid password size!")
	}

	return nil
}

func UserGetByCredentials(db *gorm.DB, email string, password string) (User, error) {
	user := UserOneByEmail(infra.S.Conn, email)

	if user.ID == 0 {
		return user, errors.Errorf("Couldn't find an user with the email '%s'!", email)
	}

	pepper := user.Pepper
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pepper+password)); err != nil {
		return user, errors.Errorf("Wrong password! (%s)", err)
	}

	return user, nil
}

func UserOneCreate(db *gorm.DB, user *User) error {
	return db.Create(&user).Error
}

func UserOneUpdate(db *gorm.DB, user *User) error {
	return db.Model(&user).Updates(user).Error
}
