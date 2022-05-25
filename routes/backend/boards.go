package backend

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"upspin.io/errors"

	"github.com/Owicca/chan/infra"
	"github.com/Owicca/chan/models/boards"
	"github.com/Owicca/chan/models/logs"
	"github.com/Owicca/chan/models/threads"
	"github.com/Owicca/chan/models/topics"
)

func init() {
	adminRouter := infra.S.Router.PathPrefix("/admin").Subrouter()
	adminRouter.HandleFunc("/boards/", http.HandlerFunc(BoardList)).Methods(http.MethodGet).Name("board_list")

	adminRouter.HandleFunc("/boards/add/", http.HandlerFunc(BoardOneAdd)).Methods(http.MethodGet).Name("board_one_add")
	adminRouter.HandleFunc("/boards/", http.HandlerFunc(BoardOneUpdate)).Methods(http.MethodPost).Name("topic_one_add_post")

	adminRouter.HandleFunc("/boards/{id:[0-9]+}/", http.HandlerFunc(BoardOne)).Methods(http.MethodGet).Name("board_one")
	adminRouter.HandleFunc("/boards/{id:[0-9]+}/", http.HandlerFunc(BoardOneUpdate)).Methods(http.MethodPost).Name("board_one_update")

	adminRouter.HandleFunc("/boards/{id:[0-9]+}/threads/", http.HandlerFunc(BoardListThreadList)).Methods(http.MethodGet).Name("board_list_thread_list")
}

func BoardList(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.BoardList"

	data := map[string]any{
		"board_list": boards.BoardListWithThreadCount(infra.S.Conn),
	}

	infra.S.HTML(w, http.StatusOK, "back/board_list", data)
}

func BoardListThreadList(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.BoardListThreadList"
	vars := mux.Vars(r)

	board_id, err := strconv.Atoi(vars["id"])
	if err != nil || board_id < 1 {
		logs.LogWarn(op, errors.Str("No id provided!"))
		infra.S.HTML(w, http.StatusOK, "back/board_list", nil)
		return
	}

	data := map[string]any{
		"thread_list": threads.BoardThreadPreviewList(infra.S.Conn, board_id),
	}

	infra.S.HTML(w, http.StatusOK, "back/board_list_thread_list", data)
}

func BoardOneAdd(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.BoardOneAdd"

	data := map[string]any{
		"board":      boards.Board{},
		"topic_list": topics.TopicList(infra.S.Conn),
	}

	infra.S.HTML(w, http.StatusOK, "back/board_one", data)
}

func BoardOne(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.BoardOne"
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil || id < 1 {
		logs.LogWarn(op, errors.Str("No id provided!"))
		infra.S.HTML(w, http.StatusOK, "back/board_one", nil)
		return
	}

	data := map[string]any{
		"board":      boards.BoardOne(infra.S.Conn, id),
		"topic_list": topics.TopicList(infra.S.Conn),
	}

	infra.S.HTML(w, http.StatusOK, "back/board_one", data)
}

func BoardOneUpdate(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.BoardOneUpdate"
	vars := mux.Vars(r)

	form_deleted_at := r.PostFormValue("deleted_at")
	var deleted_at int64 = 0
	if form_deleted_at == "1" {
		deleted_at = time.Now().Unix()
	}
	topic_id, _ := strconv.Atoi(r.PostFormValue("topic_id"))

	val, ok := vars["id"]
	the_id := 0
	if ok { // update
		id, err := strconv.Atoi(val)
		if err != nil || id < 1 {
			logs.LogWarn(op, errors.Str("No id provided!"))
			infra.S.HTML(w, http.StatusOK, "back/board_one_add", nil)
			return
		}
		the_id = id

		boards.BoardOneUpdate(infra.S.Conn, boards.Board{
			ID:          id,
			Name:        r.PostFormValue("name"),
			Topic_id:    topic_id,
			Code:        r.PostFormValue("code"),
			Description: r.PostFormValue("description"),
			Deleted_at:  deleted_at,
		})
	} else { // create
		boards.BoardOneCreate(infra.S.Conn, boards.Board{
			Name:        r.PostFormValue("name"),
			Topic_id:    topic_id,
			Code:        r.PostFormValue("code"),
			Description: r.PostFormValue("description"),
			Deleted_at:  0,
		})
	}

	data := map[string]any{
		"board":      boards.BoardOne(infra.S.Conn, the_id),
		"topic_list": topics.TopicList(infra.S.Conn),
	}

	infra.S.HTML(w, http.StatusOK, "back/board_one", data)
}
