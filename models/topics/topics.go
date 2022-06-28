package topics

import (
	"gorm.io/gorm"

	"github.com/Owicca/chan/models/boards"
)

type Topic struct {
	ID         int `gorm:"primaryKey;column:id"`
	Name       string
	Deleted_at int64
	BoardList  []boards.Board `gorm:"foreignKey:topic_id"`
}

func TopicList(db *gorm.DB) []Topic {
	var topicList []Topic

	db.Find(&topicList)

	return topicList
}

func TopicListWithBoardList(db *gorm.DB, limit int, offset int) []Topic {
	var topicList []Topic

	stmt := db.Preload("BoardList")
	if limit > 0 {
		stmt = stmt.Limit(limit)
	}
	if offset > 0 {
		stmt = stmt.Offset(offset)
	}

	stmt.Find(&topicList)

	return topicList
}

func TopicListCount(db *gorm.DB) int {
	var count int

	db.Raw(`
	SELECT COUNT(id) FROM topics
	WHERE deleted_at = 0;
	`).Scan(&count)

	return count
}

func TopicListWithBoardListWithThreadCount(db *gorm.DB) []Topic {
	var topicList []Topic

	db.Preload("BoardList.ThreadList").Find(&topicList)

	return topicList
}

func TopicOne(db *gorm.DB, id int) Topic {
	var topic Topic

	db.First(&topic, id)

	return topic
}

func TopicOneWithBoardList(db *gorm.DB, id int) Topic {
	var topic Topic

	db.Preload("BoardList.ThreadList").First(&topic, id)

	return topic
}

func TopicOneCreate(db *gorm.DB, topic *Topic) error {
	return db.Create(&topic).Error
}

func TopicOneUpdate(db *gorm.DB, topic *Topic) error {
	return db.Model(&topic).Select("*").Updates(topic).Error
}
