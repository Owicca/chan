package threads

import (
	"github.com/Owicca/chan/models/posts"
	"gorm.io/gorm"
)

type Thread struct {
	ID              int `gorm:"primaryKey;column:id"`
	Deleted_at      int64
	Board_id        int
	Primary_post_id int
	Subject         string
	Content         string       `gorm:"-"`
	Preview         []posts.Post `gorm:"foreignKey:thread_id;references:id"`
}

func BoardThreadListByCode(db *gorm.DB, board_code string) []Thread {
	var threadList []Thread

	db.Raw(`
	SELECT t.*, p.content FROM threads t
	INNER JOIN boards b ON t.board_id = b.id AND b.code = ?
	LEFT JOIN posts p ON t.primary_post_id = p.id
	`, board_code).Scan(&threadList)

	return threadList
}

func ThreadPreviewByCode(db *gorm.DB, board_code string, limit int) []Thread {
	var (
		threadList []Thread
		board_id   int
	)

	db.Raw("SELECT id FROM boards WHERE code = ?", board_code).Scan(&board_id)
	db.Preload("Preview").Find(&threadList, "board_id = ?", board_id)

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

func ThreadOneCreate(db *gorm.DB, thread *Thread) error {
	return db.Create(&thread).Error
}

func ThreadOneUpdate(db *gorm.DB, thread Thread) error {
	return db.Model(&thread).Updates(thread).Error
}
