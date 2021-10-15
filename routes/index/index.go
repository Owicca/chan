package index

import (
	"net/http"
	// "fmt"
	// "log"

	"github.com/Owicca/chan/infra"
	"github.com/Owicca/chan/models/boards"
)

func Index(srv *infra.Server) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"topics": map[string]interface{} {
				"Col 1": boards.GetBoardList(srv.Conn),
			},
			"site": map[string]interface{}{
				"name": "Chan",
				"title": "Home",
				"welcome": `Chan is a simple image-based bulletin board where anyone can post comments and share images. There are boards dedicated to a variety of topics, from Japanese animation and culture to videogames, music, and photography. Users do not need to register an account before participating in the community. Feel free to click on a board below that interests you and jump right in!

Be sure to familiarize yourself with the Rules before posting, and read the FAQ if you wish to learn more about how to use the site.`,
			},
		}

		srv.Render(w, http.StatusOK, "front/index", data)
	})
}