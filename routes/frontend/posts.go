package frontend

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Owicca/chan/infra"
	"github.com/Owicca/chan/models/logs"
	"github.com/Owicca/chan/models/posts"
	"github.com/gorilla/mux"
	"upspin.io/errors"
)

func init() {
	infra.S.HandleFunc("/boards/{board_code:[a-z0-9]+}/threads/{thread_id:[0-9]+}/", PostList).Methods(http.MethodGet).Name("f_post_list")
}

func PostList(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "front.PostList"
	vars := mux.Vars(r)

	thread_id, err := strconv.Atoi(vars["thread_id"])
	if err != nil {
		logs.LogWarn(op, errors.Str("No thread_id provided!"))
		infra.S.Redirect(w, r, http.StatusNotFound, infra.S.GenerateUrl(fmt.Sprintf("/boards/%s/threads/", vars["board_code"])))
		return
	}
	post_list := posts.ThreadPostList(infra.S.Conn, thread_id)
	data := map[string]any{
		"post_list": post_list,
	}

	infra.S.HTML(w, http.StatusOK, "front/post_list", data)
}
