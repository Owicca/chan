package frontend

import (
	"net/http"

	"github.com/Owicca/chan/internal/infra"
	"github.com/Owicca/chan/internal/models/boards"
	"github.com/Owicca/chan/internal/models/media"
	"github.com/Owicca/chan/internal/models/posts"
	"github.com/Owicca/chan/internal/models/threads"
	"github.com/Owicca/chan/internal/models/topics"
	"github.com/Owicca/chan/internal/models/users"

	"upspin.io/errors"
)

func init() {
	infra.S.HandleFunc("/", Index).Methods(http.MethodGet)
	infra.S.HandleFunc("/search/", Search).Methods(http.MethodGet)
	infra.S.HandleFunc("/search/", Search).Methods(http.MethodPost)
}

func Index(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "front.Index"

	data := map[string]any{
		"site": map[string]any{
			"name":  "Imageboard",
			"title": "Home",
			"welcome": `Imageboard is a simple image-based bulletin board where anyone can post comments and share images. There are boards dedicated to a variety of topics, from animation and culture to videogames, music, and photography. Users do not need to register an account before participating in the community. Feel free to click on a board below that interests you and jump right in!

Be sure to familiarize yourself with the Rules before posting, and read the FAQ if you wish to learn more about how to use the site.`,
		},
		"topic_list": topics.TopicListWithBoardList(infra.S.Conn, 0, 0),
		"stats": map[string]any{
			"total_posts":          posts.TotalActivePosts(infra.S.Conn),
			"total_users":          users.TotalActiveUsers(infra.S.Conn),
			"total_active_content": media.TotalMediaSize(infra.S.Conn),
		},
	}

	infra.S.HTML(w, r, http.StatusOK, "front/index", data)
}

func Search(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "front.Search"

	search := r.PostFormValue("search")
	board_code := r.PostFormValue("board_code")
	thread_id_list := posts.PostSearch(infra.S.Conn, board_code, search)

	data := map[string]any{
		"search":      search,
		"board_code":  board_code,
		"board_list":  boards.BoardList(infra.S.Conn),
		"thread_list": threads.ThreadPreviewByIdList(infra.S.Conn, thread_id_list, 0, 0),
	}

	infra.S.HTML(w, r, http.StatusOK, "front/search", data)
}
