package posts

import (
	"gorm.io/gorm"

	"github.com/Owicca/chan/models/threads"
	"github.com/Owicca/chan/models/media"
)

type Post struct {
	ID int `gorm:"primaryKey;column:post_id"`
	Deleted_at int
	Status string
	Content string
	Is_primary bool
	Thread_id int
	Media media.Media `gorm:"foreignKey:object_id"`
}

func PostList(db *gorm.DB) []Post {
	var postList []Post

	db.Preload("Thread").Preload("Media", "object_id = post_id AND object_type = posts").Find(&postList)

	return postList
}

func PostOne(db *gorm.DB, id int) Post {
	var post Post

	db.Preload("Thread").Preload("Media", "object_id = ? AND object_type = posts", id).First(&post, id)

	return post
}

func PostOneCreate(db *gorm.DB, post Post) error {
	return db.Create(&post).Error
}

func PostOneUpdate(db *gorm.DB, post Post) error {
	return db.Model(&post).Updates(post).Error
}