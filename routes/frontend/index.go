package frontend

import (
	"net/http"

	"github.com/Owicca/chan/infra"
	"github.com/Owicca/chan/models/media"
	"github.com/Owicca/chan/models/posts"
	"github.com/Owicca/chan/models/topics"
	"github.com/Owicca/chan/models/users"

	"upspin.io/errors"
)

func init() {
	infra.S.HandleFunc("/", Index).Methods(http.MethodGet).Name("site_index")
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

	infra.S.HTML(w, http.StatusOK, "front/index", data)
}
