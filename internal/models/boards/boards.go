package boards

import (
	"gorm.io/gorm"

	"github.com/Owicca/chan/internal/models/media"
	"github.com/Owicca/chan/internal/models/threads"
)

type Board struct {
	ID          int `gorm:"primaryKey;column:id"`
	Created_at  int64
	Deleted_at  int64
	Name        string
	Code        string
	Description string
	Topic_id    int
	MediaList   []media.Media `gorm:"foreignKey:object_id"`
	ThreadList  []threads.Thread
}

type BoardWithThreadCount struct {
	Board
	Thread_count int
}

func BoardListWithThreadCount(db *gorm.DB, limit int, offset int) []BoardWithThreadCount {
	var boards []BoardWithThreadCount

	stmt := db.Table("boards AS b").Select("b.*, COUNT(t.id) as thread_count").Joins("LEFT JOIN threads AS t ON t.board_id=b.id")
	stmt = stmt.Where("b.deleted_at = 0").Group("b.id")
	if limit > 0 {
		stmt = stmt.Limit(limit)
	}
	if offset > 0 {
		stmt = stmt.Offset(offset)
	}

	stmt.Find(&boards)

	return boards
}

func BoardIdByCode(db *gorm.DB, board_code string) int {
	var board_id int

	db.Raw(`
	SELECT id FROM boards
	WHERE code = ?
	`, board_code).Scan(&board_id)

	return board_id
}

func BoardList(db *gorm.DB) []Board {
	var boards []Board

	db.Preload("MediaList", "object_type = 'boards'").Find(&boards)

	return boards
}

func BoardListCountOfTopic(db *gorm.DB, topic_id int) int {
	var count int

	db.Raw(`
	SELECT COUNT(id) FROM boards
	WHERE topic_id = ?
	`, topic_id).Scan(&count)

	return count
}

func BoardListCount(db *gorm.DB) int {
	var count int

	db.Raw(`
	SELECT COUNT(id) FROM boards
	WHERE deleted_at = 0;
	`).Scan(&count)

	return count
}

func BoardOne(db *gorm.DB, id int) Board {
	var board Board

	db.Preload("MediaList").First(&board, id)

	return board
}

func BoardOneCreate(db *gorm.DB, board *Board) error {
	return db.Create(&board).Error
}

func BoardOneUpdate(db *gorm.DB, board *Board) error {
	return db.Model(&board).Select("*").Updates(board).Error
}
