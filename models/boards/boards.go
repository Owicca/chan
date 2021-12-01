package boards

import (
	"gorm.io/gorm"

	"github.com/Owicca/chan/models/media"
)

type Board struct {
	ID int `gorm:"primaryKey;column:board_id"`
	Deleted_at int
	Name string
	Code string
	Description string
	MediaList []media.Media `gorm:"foreignKey:object_id"`
}

type BoardWithThreadCount struct{
	Board
	Thread_count int
}

func BoardListWithThreadCount(db *gorm.DB) []BoardWithThreadCount {
	var boards []BoardWithThreadCount

	db.Table("boards").Select("boards.*", "COUNT(t.thread_id) AS thread_count").Joins("LEFT JOIN threads AS t ON t.board_id=boards.board_id").Preload("MediaList", "object_type = 'boards'").Find(&boards)

	return boards
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
	return db.Model(&board).Updates(board).Error
}