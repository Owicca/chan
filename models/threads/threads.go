package threads

import (
	"gorm.io/gorm"
)

type Thread struct {
	ID int `gorm:"primaryKey;column:thread_id"`
	Deleted_at int64
	Board_id int
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