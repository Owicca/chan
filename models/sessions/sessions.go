package sessions

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/Owicca/chan/models/logs"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

type Session struct {
	ID   int `gorm:"primaryKey;column:id"`
	Data string
}

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

func Get(db *gorm.DB, session_id int) Session {
	var session Session

	db.First(&session, session_id)

	return session
}

func Update(db *gorm.DB, session_id int, data map[string]any) {
	d := []byte{}

	d, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("err while unmarshaling (%s)", err)
	}

	db.Exec(`
		REPLACE INTO sessions VALUES(?,?)
	`, session_id, d)
}

func Delete(db *gorm.DB, session_id int) {
	db.Delete(&Session{}, session_id)
}
