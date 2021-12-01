package backend

import (
	"github.com/Owicca/chan/infra"
	"net/http"

	"upspin.io/errors"
	"github.com/gorilla/mux"
	"strconv"
	"time"
	
	"github.com/Owicca/chan/models/logs"
	"github.com/Owicca/chan/models/threads"
	"github.com/Owicca/chan/models/boards"
)

func init() {
	adminRouter := infra.S.Router.PathPrefix("/admin").Subrouter()
	adminRouter.HandleFunc("/threads/", http.HandlerFunc(ThreadList)).Methods(http.MethodGet).Name("thread_list")
	adminRouter.HandleFunc("/boards/{board_id:[0-9]+}/threads/", http.HandlerFunc(BoardThreadList)).Methods(http.MethodGet).Name("board_thread_list")
	adminRouter.HandleFunc("/threads/{thread_id:[0-9]+}/", http.HandlerFunc(ThreadOne)).Methods(http.MethodGet).Name("thread_one")
	adminRouter.HandleFunc("/threads/{thread_id:[0-9]+}/", http.HandlerFunc(ThreadOneUpdate)).Methods(http.MethodPost).Name("thread_one_update")
}

func BoardThreadList(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.BoardThreadList"
	vars := mux.Vars(r)

	board_id, err := strconv.Atoi(vars["board_id"])
	if err != nil || board_id < 1 {
		logs.LogWarn(op, errors.Str("No board_id provided!"))
		infra.S.Redirect(w, r, http.StatusNotFound, infra.S.GenerateUrl("/admin/boards/"))
		return
	}

	data := map[string]interface{} {
		"threads": threads.ThreadList(infra.S.Conn),
	}

	infra.S.HTML(w, http.StatusOK, "back/thread_list", data)
}

func ThreadList(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.ThreadList"

	data := map[string]interface{} {
		"threads": threads.ThreadList(infra.S.Conn),
	}

	infra.S.HTML(w, http.StatusOK, "back/thread_list", data)
}

func ThreadOne(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.ThreadOne"
	vars := mux.Vars(r)

	thread_id, err := strconv.Atoi(vars["thread_id"])
	if err != nil || thread_id < 1 {
		logs.LogWarn(op, errors.Str("No thread_id provided!"))
		infra.S.Redirect(w, r, http.StatusNotFound, infra.S.GenerateUrl("/admin/threads/"))
		return
	}

	data := map[string]interface{} {
		"thread": threads.ThreadOne(infra.S.Conn, thread_id),
		"boardList": boards.BoardList(infra.S.Conn),
	}

	infra.S.HTML(w, http.StatusOK, "back/thread_one", data)
}

func ThreadOneUpdate(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.ThreadOneUpdate"
	if err := r.ParseForm(); err != nil {
		logs.LogErr(op, err)
		infra.S.HTML(w, http.StatusOK, "back/thread_one", nil)
		return
	}

	thread_id, err := strconv.Atoi(r.PostFormValue("thread_id"))
	if err != nil || thread_id < 1 {
		logs.LogWarn(op, errors.Str("No thread_id provided!"))
		infra.S.HTML(w, http.StatusOK, "back/thread_one", nil)
		return
	}

	board_id, err := strconv.Atoi(r.PostFormValue("board_id"))
	if err != nil || board_id < 1 {
		logs.LogWarn(op, errors.Str("No board_id provided!"))
		infra.S.HTML(w, http.StatusOK, "back/thread_one", nil)
		return
	}

	var deleted_at int64 = 0
	if r.PostFormValue("deleted_at") != "1" {
		deleted_at = time.Now().Unix()
	}

	newThread := threads.Thread{
		ID: thread_id,
		Deleted_at: deleted_at,
		Board_id: board_id,
	}

	threads.ThreadOneUpdate(infra.S.Conn, newThread)

	data := map[string]interface{} {
		"thread": threads.ThreadOne(infra.S.Conn, thread_id),
		"boardList": boards.BoardList(infra.S.Conn),
	}

	infra.S.HTML(w, http.StatusOK, "back/thread_one", data)
}