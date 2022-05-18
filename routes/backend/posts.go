package backend

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
	adminRouter := infra.S.Router.PathPrefix("/admin/").Subrouter()
	adminRouter.HandleFunc("/threads/{thread_id:[0-9]+}/posts/", http.HandlerFunc(ThreadPostList)).Methods(http.MethodGet).Name("post_list")
	adminRouter.HandleFunc("/posts/{post_id:[0-9]+}/", http.HandlerFunc(PostOne)).Methods(http.MethodGet).Name("post_one")
	adminRouter.HandleFunc("/posts/{post_id:[0-9]+}/", http.HandlerFunc(PostOneDelete)).Methods(http.MethodPost).Name("post_one_delete")
}

func ThreadPostList(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.ThreadPostList"
	vars := mux.Vars(r)

	thread_id, err := strconv.Atoi(vars["thread_id"])
	if err != nil {
		logs.LogWarn(op, errors.Str("No thread_id provided!"))
		infra.S.Redirect(w, r, "/admin/threads/")
		return
	}
	data := map[string]any{
		"posts":          posts.ThreadPostList(infra.S.Conn, thread_id),
		"postStatusList": posts.PostStatusList(),
	}

	infra.S.HTML(w, http.StatusOK, "back/post_list", data)
}

func PostOne(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.PostOne"
	vars := mux.Vars(r)

	post_id, err := strconv.Atoi(vars["post_id"])
	if err != nil {
		logs.LogWarn(op, errors.Str("No post_id provided!"))
		infra.S.Redirect(w, r, "/admin/threads/")
		return
	}
	data := map[string]any{
		"post": posts.PostOne(infra.S.Conn, post_id),
	}

	infra.S.HTML(w, http.StatusOK, "back/post_one", data)
}

func PostOneDelete(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.PostOneDelete"
	r.ParseForm()

	post_id, thread_id := r.PostFormValue("post_id"), r.PostFormValue("thread_id")
	post_id_int, _ := strconv.Atoi(post_id)
	//if err != nil {
	//	logs.LogWarn(op, errors.Str("No post_id provided!"))
	//	infra.S.Redirect(w, r, "/admin/posts/")
	//	return
	//}
	posts.PostOneDelete(infra.S.Conn, post_id_int)

	infra.S.Redirect(w, r, fmt.Sprintf("/admin/threads/%s/posts/", thread_id))
}
