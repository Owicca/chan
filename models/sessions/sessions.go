package sessions

import (
	"math/rand"
	"time"

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

const (
	charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var (
	PublicUrl = map[string]string{
		"front_index": "/",
		"login":       "/admin/login/",
		"logout":      "/admin/logout/",
	}
	seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
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

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}

func GeneratePepper(length int) string {
	return StringWithCharset(length, charset)
}
