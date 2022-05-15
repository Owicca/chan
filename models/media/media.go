package media

import "gorm.io/gorm"

type Media struct {
	Object_id   int
	Object_type string
	Deleted_at  int
	Name        string
	Code        string
	Path        string
	Thumb       string
	X           int
	Y           int
	Size        int64
}

func TotalMediaSize(db *gorm.DB) int64 {
	var count int64

	db.Raw(`
	SELECT SUM(m.size) FROM media m
	WHERE
	m.deleted_at = 0
	`).Scan(&count)

	return count
}
