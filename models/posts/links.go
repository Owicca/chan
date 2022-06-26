package posts

import "gorm.io/gorm"

type Link struct {
	Src  int
	Dest int
}

func LinkOneCreate(db *gorm.DB, l *Link) error {
	return db.Create(&l).Error
}
