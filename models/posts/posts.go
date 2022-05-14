package posts

import (
	"time"

	"gorm.io/gorm"

	"github.com/Owicca/chan/models/media"
)

type Post struct {
	ID         int `gorm:"primaryKey;column:id"`
	Deleted_at int64
	Status     string
	Thread_id  int
	Content    string
	Media      media.Media `gorm:"foreignKey:object_id"`
}

func ThreadPostList(db *gorm.DB, thread_id int) []Post {
	var postList []Post

	db.Preload("Thread").Preload("Media", "media.id = posts.id AND media.object_type = 'posts'").Find(&postList, "thread.id = ?", thread_id)

	return postList
}

func PostList(db *gorm.DB) []Post {
	var postList []Post

	db.Preload("Thread").Preload("Media", "media.id = posts.id AND media.object_type = 'posts'").Find(&postList)

	return postList
}

func PostOne(db *gorm.DB, id int) Post {
	var post Post

	db.Preload("Thread").Preload("Media", "media.id = ? AND media.object_type = 'posts'", id).First(&post, id)

	return post
}

func PostOneCreate(db *gorm.DB, post Post) error {
	return db.Create(&post).Error
}

func PostOneUpdate(db *gorm.DB, id int, post Post) error {
	return db.Model(&post).Updates(post).Error
}

func PostOneDelete(db *gorm.DB, id int) error {
	post := &Post{
		ID: id,
	}

	return db.Model(post).Updates(&Post{Deleted_at: time.Now().Unix()}).Error
}
