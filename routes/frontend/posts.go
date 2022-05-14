package frontend

import (
	"net/http"

	"github.com/Owicca/chan/infra"
	"upspin.io/errors"
)

func init() {
	infra.S.HandleFunc("/boards/{board_code:[a-z]+}/{thread_id:{0-9}+}", http.HandlerFunc(PostList)).Methods(http.MethodGet).Name("board_list")
	// infra.S.HandleFunc("/boards/{board_code:[a-z]+}/thread/{thread_id:[0-9]+}/", http.HandlerFunc(Thread)).Methods(http.MethodGet).Name("thread_one")
}

type viewPost struct {
	Name    string
	Content string
}

func PostList(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "front.PostList"

	data := map[string]any{
		"posts": []viewPost{
			{
				Name:    "p1",
				Content: "p1",
			},
			{
				Name:    "p2",
				Content: "p2",
			},
		},
	}

	infra.S.HTML(w, http.StatusOK, "front/post_list", data)
}
