package boards

import (
	// "context"

	"github.com/Owicca/chan/models/media"
	// "github.com/Owicca/chan/models/base"
	// "github.com/jackc/pgx/v4"
	// "github.com/jackc/pgx/v4/pgxpool"
)

type Board struct {
	Board_id int
	Deleted_at int
	Name string
	Code string
	Description string
	Media media.Media
}

func GetBoardList() []Board {
	boards := []Board{}

	return boards
}

// func GetOneBoard(db *pgxpool.Pool, int id) (Board, error) {
// 	sql := `SELECT b.*, mt.code AS media_type, m.path AS media_path
// 	FROM boards b
// 	LEFT JOIN medias AS m ON b.media_id=m.media_id
// 	LEFT JOIN media_types AS mt ON m.media_type_id = mt.media_type_id
// 	WHERE board_id = $1
// 	LIMIT 1`
// 	row := db.QueryRowx(context.Background(), sql, )

// 	board := &Board{}

// 	if err := db.QueryRowx(board, sql, id); err != nil {
// 		log.Fatal(err)
// 	}
// }