package frontend

import (
	"net/http"

	"github.com/Owicca/chan/infra"
	"github.com/Owicca/chan/models/topics"

	"upspin.io/errors"
)

func init() {
	infra.S.HandleFunc("/", http.HandlerFunc(Index)).Methods(http.MethodGet).Name("site_index")
}

func Index(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "front.Index"

	data := map[string]any{
		"topic_list": topics.TopicListWithBoardList(infra.S.Conn),
		"site": map[string]any{
			"name":  "Chan",
			"title": "Home",
			"welcome": `Chan is a simple image-based bulletin board where anyone can post comments and share images. There are boards dedicated to a variety of topics, from Japanese animation and culture to videogames, music, and photography. Users do not need to register an account before participating in the community. Feel free to click on a board below that interests you and jump right in!

Be sure to familiarize yourself with the Rules before posting, and read the FAQ if you wish to learn more about how to use the site.`,
		},
	}

	infra.S.HTML(w, http.StatusOK, "front/index", data)
}
