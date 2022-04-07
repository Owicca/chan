package backend

import (
	"net/http"
	"strconv"
	"log"

	"github.com/Owicca/chan/infra"
	"github.com/Owicca/chan/models/posts"
	"github.com/Owicca/chan/models/logs"

	"github.com/gorilla/mux"
	"upspin.io/errors"
)

func init() {
	adminRouter := infra.S.Router.PathPrefix("/admin/").Subrouter()
	adminRouter.HandleFunc("/threads/{thread_id:[0-9]+}/posts/", http.HandlerFunc(ThreadPostList)).Methods(http.MethodGet).Name("post_list")
	adminRouter.HandleFunc("/posts/{post_id:[0-9]+}/", http.HandlerFunc(PostOne)).Methods(http.MethodGet).Name("post_one")
	adminRouter.HandleFunc("/posts/{post_id:[0-9]+}/", http.HandlerFunc(PostOneDelete)).Methods(http.MethodDelete).Name("post_one")
}

func ThreadPostList(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.ThreadPostList"
	vars := mux.Vars(r)

	thread_id, err := strconv.Atoi(vars["thread_id"])
	if err != nil {
		logs.LogWarn(op, errors.Str("No thread_id provided!"))
		infra.S.Redirect(w, r, http.StatusNotFound, infra.S.GenerateUrl("/admin/threads/"))
		return
	}
	data := map[string]interface{} {
		"posts": posts.ThreadPostList(infra.S.Conn, thread_id),
		"postStatusList": posts.PostStatusList(),
	}
	log.Println("list => ", posts.ThreadPostList(infra.S.Conn, thread_id))

	infra.S.HTML(w, http.StatusOK, "back/post_list", data)
}

func PostOne(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.PostOne"
	vars := mux.Vars(r)

	post_id, err := strconv.Atoi(vars["post_id"])
	if err != nil {
		logs.LogWarn(op, errors.Str("No post_id provided!"))
		infra.S.Redirect(w, r, http.StatusNotFound, infra.S.GenerateUrl("/admin/threads/"))
		return
	}
	data := map[string]interface{} {
		"post": posts.PostOne(infra.S.Conn, post_id),
	}

	infra.S.HTML(w, http.StatusOK, "back/post_one", data)
}

func PostOneDelete(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.PostOneDelete"
	vars := mux.Vars(r)

	post_id, err := strconv.Atoi(vars["post_id"])
	if err != nil {
		logs.LogWarn(op, errors.Str("No post_id provided!"))
		infra.S.Redirect(w, r, http.StatusNotFound, infra.S.GenerateUrl("/admin/posts/"))
		return
	}
	posts.PostOneDelete(infra.S.Conn, post_id)

	infra.S.Redirect(w, r, http.StatusOK, infra.S.GenerateUrl("/admin/posts/"))
}