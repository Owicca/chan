package board

import (
	"net/http"
	// "fmt"
	// "log"

	"github.com/Owicca/chan/infra"
	// "github.com/Owicca/chan/models/boards"
)

func Board(srv *infra.Server) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{} {
			"threads": []struct{
				Name string
				Path string
			} {
				{
					Name: "n1",
					Path: "/a/n1/",
				},
				{
					Name: "n2",
					Path: "/a/n2/",
				},
			},
		}

		// board := boards.GetOneBoard(srv.Conn, 1)


		srv.Render(w, http.StatusOK, "front/board_index", data)
	})
}