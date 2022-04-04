package topics

import (
	"gorm.io/gorm"

	"github.com/Owicca/chan/models/boards"
)

type Topic struct {
	ID int `gorm:"primaryKey;column:id"`
	Name string
	Deleted_at int64
	BoardList []boards.Board `gorm:"foreignKey:topic_id"`
}

func TopicList(db *gorm.DB) []Topic {
	var topicList []Topic

	db.Find(&topicList)

	return topicList
}

func TopicListWithBoardList(db *gorm.DB) []Topic {
	var topicList []Topic

	db.Preload("BoardList").Find(&topicList)

	return topicList
}

func TopicOne(db *gorm.DB, id int) Topic {
	var topic Topic

	db.First(&topic, id)

	return topic
}

func TopicOneCreate(db *gorm.DB, topic Topic) error {
	return db.Create(&topic).Error
}

func TopicOneUpdate(db *gorm.DB, topic Topic) error {
	return db.Model(&topic).Select("*").Updates(topic).Error
}