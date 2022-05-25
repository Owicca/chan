package boards

import (
	"gorm.io/gorm"

	"github.com/Owicca/chan/models/media"
	"github.com/Owicca/chan/models/threads"
)

type Board struct {
	ID          int `gorm:"primaryKey;column:id"`
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

func BoardListWithThreadCount(db *gorm.DB) []BoardWithThreadCount {
	var boards []BoardWithThreadCount

	db.Preload("MediaList", "object_type = 'boards'").Raw(`
		SELECT b.*, COUNT(t.id) as thread_count FROM boards AS b
		LEFT JOIN threads AS t ON t.board_id=b.id
		GROUP BY b.id
	`).Find(&boards)

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

func BoardOne(db *gorm.DB, id int) Board {
	var board Board

	db.Preload("MediaList").First(&board, id)

	return board
}

func BoardOneCreate(db *gorm.DB, board Board) error {
	return db.Create(&board).Error
}

func BoardOneUpdate(db *gorm.DB, board Board) error {
	return db.Model(&board).Select("*").Updates(board).Error
}
