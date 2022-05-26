package sessions

import (
	"github.com/Owicca/chan/models/logs"
	"golang.org/x/crypto/bcrypt"
	"upspin.io/errors"
)

//func GeneratePassword(password string, salt string) string {
//	iter := 4096
//	keyLen := 32
//
//	hash := pbkdf2.Key([]byte(password), []byte(salt), iter, keyLen, sha512.New512_256)
//
//	return hex.EncodeToString(hash[:])
//}

var (
	PublicUrl = map[string]string{
		"login":  "/admin/login/",
		"logout": "/admin/logout/",
	}
)

func GeneratePassword(password string, pepper string) string {
	const op errors.Op = "sessions.GeneratePassword"
	results := ""
	cost := bcrypt.DefaultCost

	hash, err := bcrypt.GenerateFromPassword([]byte(pepper+password), cost)
	if err != nil {
		logs.LogErr(op, errors.Errorf("Couldn't hash password '%s' (%s)!", password, err))

		return results
	}
	results = string(hash)

	return results
}
