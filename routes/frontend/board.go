package frontend

import (
	"net/http"

	"github.com/Owicca/chan/infra"
	"upspin.io/errors"
)

func init() {
	infra.S.HandleFunc("/{board_name:[a-z]+}/", http.HandlerFunc(Board)).Methods(http.MethodGet).Name("board_index")
}

type viewThread struct {
	Name string
	Path string
}

func Board(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "front.Board"
	println(op)
	data := map[string]interface{} {
		"threads": []viewThread{
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


	infra.S.HTML(w, http.StatusOK, "front/board_index", data)
}