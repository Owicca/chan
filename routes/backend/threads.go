package backend

import (
	"net/http"

	"github.com/Owicca/chan/infra"

	"strconv"
	"time"

	"github.com/gorilla/mux"
	"upspin.io/errors"

	"github.com/Owicca/chan/models/boards"
	"github.com/Owicca/chan/models/logs"
	"github.com/Owicca/chan/models/posts"
	"github.com/Owicca/chan/models/threads"
)

func init() {
	adminRouter := infra.S.Router.PathPrefix("/admin").Subrouter()
	adminRouter.HandleFunc("/threads/", ThreadList).Methods(http.MethodGet).Name("thread_list")
	adminRouter.HandleFunc("/threads/p{page:[0-9]+}/", ThreadList).Methods(http.MethodGet).Name("thread_list_page")
	adminRouter.HandleFunc("/boards/{board_id:[0-9]+}/threads/", BoardThreadList).Methods(http.MethodGet).Name("board_thread_list")
	adminRouter.HandleFunc("/threads/{thread_id:[0-9]+}/", ThreadOne).Methods(http.MethodGet).Name("thread_one")
	adminRouter.HandleFunc("/threads/{thread_id:[0-9]+}/", ThreadOneUpdate).Methods(http.MethodPost).Name("thread_one_update")
}

func BoardThreadList(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.BoardThreadList"
	vars := mux.Vars(r)

	board_id, err := strconv.Atoi(vars["board_id"])
	if err != nil || board_id < 1 {
		logs.LogWarn(op, errors.Str("No board_id provided!"))
		infra.S.Redirect(w, r, infra.S.GenerateUrl("/admin/boards/"))
		return
	}

	data := map[string]any{
		"threads": threads.ThreadList(infra.S.Conn),
	}

	infra.S.HTML(w, r, http.StatusOK, "back/thread_list", data)
}

func ThreadList(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.ThreadList"
	vars := mux.Vars(r)

	//page_limit := int(infra.S.Data["page_limit"].(float64))
	page, _ := strconv.Atoi(vars["page"])
	offset := page * posts.PostPageLimit
	totalThreads := threads.ThreadPreviewListCount(infra.S.Conn)
	threads := threads.ThreadPreviewList(infra.S.Conn, threads.ThreadPageLimit, offset)
	pageCount, pageHelper := infra.GeneratePagination(totalThreads, posts.PostPageLimit)

	data := map[string]any{
		"threads":     threads,
		"page_count":  pageCount,
		"page_helper": pageHelper,
		"page":        page,
		"base_url":    "/admin/threads/",
	}

	infra.S.HTML(w, r, http.StatusOK, "back/thread_list", data)
}

func ThreadOne(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.ThreadOne"
	vars := mux.Vars(r)

	thread_id, err := strconv.Atoi(vars["thread_id"])
	if err != nil || thread_id < 1 {
		logs.LogWarn(op, errors.Str("No thread_id provided!"))
		infra.S.Redirect(w, r, infra.S.GenerateUrl("/admin/threads/"))
		return
	}

	data := map[string]any{
		"thread":    threads.ThreadOne(infra.S.Conn, thread_id),
		"boardList": boards.BoardList(infra.S.Conn),
	}

	infra.S.HTML(w, r, http.StatusOK, "back/thread_one", data)
}

func ThreadOneUpdate(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.ThreadOneUpdate"
	if err := r.ParseForm(); err != nil {
		logs.LogErr(op, err)
		infra.S.HTML(w, r, http.StatusOK, "back/thread_one", nil)
		return
	}

	thread_id, err := strconv.Atoi(r.PostFormValue("thread_id"))
	if err != nil || thread_id < 1 {
		logs.LogWarn(op, errors.Str("No thread_id provided!"))
		infra.S.HTML(w, r, http.StatusOK, "back/thread_one", nil)
		return
	}

	board_id, err := strconv.Atoi(r.PostFormValue("board_id"))
	if err != nil || board_id < 1 {
		logs.LogWarn(op, errors.Str("No board_id provided!"))
		infra.S.HTML(w, r, http.StatusOK, "back/thread_one", nil)
		return
	}

	var deleted_at int64 = 0
	if r.PostFormValue("deleted_at") != "1" {
		deleted_at = time.Now().Unix()
	}

	newThread := threads.Thread{
		ID:         thread_id,
		Deleted_at: deleted_at,
		Board_id:   board_id,
	}

	threads.ThreadOneUpdate(infra.S.Conn, newThread)

	data := map[string]any{
		"thread":    threads.ThreadOne(infra.S.Conn, thread_id),
		"boardList": boards.BoardList(infra.S.Conn),
	}

	infra.S.HTML(w, r, http.StatusOK, "back/thread_one", data)
}
