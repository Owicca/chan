package posts

import "gorm.io/gorm"

type Backlink struct {
	ID      int
	Post_id int
	Link    int
}

func BacklinkOneCreate(db *gorm.DB, bl *Backlink) error {
	return db.Create(&bl).Error
}
