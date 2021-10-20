package frontend

import (
	"net/http"

	"github.com/Owicca/chan/infra"
	"upspin.io/errors"
)

func init() {
	infra.S.HandleFunc("/{board_name:[a-z]+}/thread/{thread_id:[0-9]+}/", http.HandlerFunc(Thread)).Methods(http.MethodGet).Name("thread_index")
}

type viewPost struct {
	Name string
	Content string
}

func Thread(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "front.Thread"
	println(op)
	data := map[string]interface{} {
		"posts": []viewPost{
			{
				Name: "p1",
				Content: "p1",
			},
			{
				Name: "p2",
				Content: "p2",
			},
		},
	}

	infra.S.HTML(w, http.StatusOK, "front/thread_index", data)
}