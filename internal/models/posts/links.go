package posts

import "gorm.io/gorm"

type Link struct {
	Src  int
	Dest int
}

func LinkOneCreate(db *gorm.DB, l *Link) error {
	return db.Create(&l).Error
}

func LinkListCreate(db *gorm.DB, lSrc int, lDest []int) error {
	var links []*Link

	for d := range lDest {
		links = append(links, &Link{
			Src:  lSrc,
			Dest: d,
		})
	}

	return db.Create(&links).Error
}
