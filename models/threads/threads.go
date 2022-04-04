package threads

import (
	"gorm.io/gorm"

	"github.com/Owicca/chan/models/posts"
)

type Thread struct {
	ID int `gorm:"primaryKey;column:id"`
	Deleted_at int64
	Board_id int
	Primary posts.Post `gorm:"foreignKey:id"`
}

func BoardThreadListByCode(db *gorm.DB, board_code string) []Thread {
	var threadList []Thread

	db.Joins("INNER JOIN boards ON threads.board_id = boards.id && boards.code = ?", board_code).Preload("Primary", "is_primary = ?", 1).Find(&threadList)

	return threadList
}

func BoardThreadList(db *gorm.DB, board_id int) []Thread {
	var threadList []Thread

	db.Find(&threadList, "board_id = ?", board_id)

	return threadList
}

func ThreadList(db *gorm.DB) []Thread {
	threadList := []Thread{}

	db.Find(&threadList)

	return threadList
}

func ThreadOne(db *gorm.DB, id int) Thread {
	thread := Thread{}

	db.First(&thread, id)

	return thread
}

func ThreadOneCreate(db *gorm.DB, thread Thread) error {
	return db.Create(&thread).Error
}

func ThreadOneUpdate(db *gorm.DB, thread Thread) error {
	return db.Model(&thread).Updates(thread).Error
}