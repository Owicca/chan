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

func BoardList(db *gorm.DB) []Board {
	var boards []Board

	db.Preload("MediaList").Find(&boards)

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