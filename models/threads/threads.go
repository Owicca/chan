package threads

import (
	"github.com/Owicca/chan/models/posts"
	"gorm.io/gorm"
)

const (
	ThreadPageLimit = 15
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

func ThreadPreviewByCode(db *gorm.DB, board_code string, limit int, offset int) []Thread {
	var (
		threadList []Thread
		board_id   int
	)

	db.Raw("SELECT id FROM boards WHERE code = ?", board_code).Scan(&board_id)
	stmt := db.Preload("Preview").Preload("Preview.LinkList")
	if limit > 0 {
		stmt = stmt.Limit(limit)
	}
	if offset > 0 {
		stmt = stmt.Offset(offset)
	}

	stmt.Find(&threadList, "board_id = ?", board_id)

	return threadList
}

func ThreadPreviewByIdList(db *gorm.DB, thread_id_list []int, limit int, offset int) []Thread {
	var (
		threadList []Thread
	)

	stmt := db.Preload("Preview").Preload("Preview.LinkList")
	if limit > 0 {
		stmt = stmt.Limit(limit)
	}
	if offset > 0 {
		stmt = stmt.Offset(offset)
	}

	stmt.Find(&threadList, "id IN (?)", thread_id_list)

	return threadList
}

func TotalActiveThreads(db *gorm.DB, board_code string) int {
	var count int

	db.Raw(`
	SELECT COUNT(t.id) FROM threads t
	LEFT JOIN boards b ON b.id = t.board_id
	WHERE
	b.code = ?
	AND t.deleted_at = ?
	`, board_code, 0).Scan(&count)

	return count
}

func BoardThreadList(db *gorm.DB, board_id int) []Thread {
	var threadList []Thread

	db.Find(&threadList, "board_id = ?", board_id)

	return threadList
}

func ThreadListCountOfBoard(db *gorm.DB, board_id int) int {
	var count int

	db.Raw(`
	SELECT COUNT(id) FROM threads
	WHERE board_id = ?
	`, board_id).Scan(&count)

	return count
}

func BoardThreadPreviewList(db *gorm.DB, board_id int, limit int, offset int) []Thread {
	var threadList []Thread

	stmt := db.Preload("Preview")
	if limit > 0 {
		stmt = stmt.Limit(limit)
	}
	if offset > 0 {
		stmt = stmt.Offset(offset)
	}

	stmt.Find(&threadList, "board_id = ?", board_id)

	return threadList
}

func ThreadList(db *gorm.DB) []Thread {
	threadList := []Thread{}

	db.Find(&threadList)

	return threadList
}

func ThreadListByIdList(db *gorm.DB, id_list []int) []Thread {
	var threadList []Thread

	if len(id_list) > 0 {
		db.Find(&threadList, id_list)
	}

	return threadList
}

func ThreadPreviewListCount(db *gorm.DB) int {
	var count int

	db.Raw(`
	SELECT COUNT(id) FROM threads
	WHERE deleted_at = 0;
	`).Scan(&count)

	return count
}

func ThreadPreviewList(db *gorm.DB, limit int, offset int) []Thread {
	var threadList []Thread

	stmt := db.Preload("Preview")
	if limit > 0 {
		stmt = stmt.Limit(limit)
	}
	if offset > 0 {
		stmt = stmt.Offset(offset)
	}

	stmt.Find(&threadList)

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
