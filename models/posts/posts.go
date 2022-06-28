package posts

import (
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/Owicca/chan/models/media"
	"github.com/Owicca/chan/models/tripkeys"
)

const (
	PostPageLimit = 15
)

type Post struct {
	ID             int `gorm:"primaryKey;column:id"`
	Created_at     int64
	Deleted_at     int64
	Tripcode       string
	SecureTripcode string
	Status         string
	Thread_id      int
	Name           string
	Content        string
	Media          media.Media `gorm:"foreignKey:object_id;references:id"`
	LinkList       []Link      `gorm:"foreignKey:src;references:id"`
}

func ThreadPostList(db *gorm.DB, thread_id int, limit int, offset int) []Post {
	var postList []Post

	stmt := db.Preload("Media", "media.object_type = 'posts'").Preload("LinkList")
	if limit > 0 {
		stmt = stmt.Limit(limit)
	}
	if offset > 0 {
		stmt = stmt.Offset(offset)
	}

	stmt.Find(&postList, "thread_id = ?", thread_id)

	return postList
}

func TotalActivePosts(db *gorm.DB) int {
	var count int

	db.Raw(`
	SELECT COUNT(p.id) FROM posts p
	WHERE
	p.deleted_at = ?
	AND p.status = ?
	`, 0, PostStatusActive).Scan(&count)

	return count
}

func PostList(db *gorm.DB) []Post {
	var postList []Post

	db.Preload("Thread").Preload("Media", "media.object_id = posts.id AND media.object_type = 'posts'").Find(&postList)

	return postList
}

func PostOne(db *gorm.DB, id int) Post {
	var post Post

	db.Preload("Media", "media.object_id = ? AND media.object_type = 'posts'", id).First(&post, id)

	return post
}

func PostOneCreate(db *gorm.DB, post *Post) error {
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

func DeconstructInput(inp string) (string, string, string) {
	var (
		name   string
		trip   string
		secure string
	)

	if inp != "" {
		//log.Println("not empty")
		parts := strings.Split(inp, "##")
		if len(parts) > 1 {
			//log.Println("secure", parts)
			name = parts[0]
			if parts[1] != "" {
				secure = tripkeys.Tripkey([]byte(parts[1]))
			}
			secure = parts[1]
		} else {
			parts = strings.Split(inp, "#")
			//log.Print("trip parts", parts)
			if len(parts) > 1 {
				//log.Println("trip", parts)
				name = parts[0]
				if parts[1] != "" {
					trip = tripkeys.Tripkey([]byte(parts[1]))
				}
			}
		}
	}

	//log.Println("deconstructed", name, trip, secure)
	return name, trip, secure
}
