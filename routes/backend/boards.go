package backend

import (
	"fmt"
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
	adminRouter.HandleFunc("/boards/", BoardList).Methods(http.MethodGet).Name("board_list")
	adminRouter.HandleFunc("/boards/p{page:[0-9]+}/", BoardList).Methods(http.MethodGet).Name("board_list_page")

	adminRouter.HandleFunc("/boards/add/", BoardOneAdd).Methods(http.MethodGet).Name("board_one_add")
	adminRouter.HandleFunc("/boards/", BoardOneCreate).Methods(http.MethodPost).Name("board_one_create")

	adminRouter.HandleFunc("/boards/{id:[0-9]+}/", BoardOne).Methods(http.MethodGet).Name("board_one")
	adminRouter.HandleFunc("/boards/{id:[0-9]+}/", BoardOneUpdate).Methods(http.MethodPost).Name("board_one_update")

	adminRouter.HandleFunc("/boards/{id:[0-9]+}/threads/", BoardListThreadList).Methods(http.MethodGet).Name("board_list_thread_list")
	adminRouter.HandleFunc("/boards/{id:[0-9]+}/threads/p{page:[0-9]+}/", BoardListThreadList).Methods(http.MethodGet).Name("board_list_thread_list_page")
}

func BoardList(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.BoardList"
	vars := mux.Vars(r)

	page_limit := int(infra.S.Data["page_limit"].(float64))
	page, _ := strconv.Atoi(vars["page"])
	offset := page * page_limit
	total := boards.BoardListCount(infra.S.Conn)
	board_list := boards.BoardListWithThreadCount(infra.S.Conn, page_limit, offset)
	pageCount, pageHelper := infra.GeneratePagination(total, page_limit)

	data := map[string]any{
		"board_list":  board_list,
		"page_count":  pageCount,
		"page_helper": pageHelper,
		"page":        page,
		"base_url":    "/admin/boards/",
	}

	infra.S.HTML(w, r, http.StatusOK, "back/board_list", data)
}

func BoardListThreadList(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.BoardListThreadList"
	vars := mux.Vars(r)

	board_id, err := strconv.Atoi(vars["id"])
	if err != nil || board_id < 1 {
		logs.LogWarn(op, errors.Str("No id provided!"))
		infra.S.HTML(w, r, http.StatusOK, "back/board_list", nil)
		return
	}

	page_limit := int(infra.S.Data["page_limit"].(float64))
	page, _ := strconv.Atoi(vars["page"])
	offset := page * page_limit
	total := threads.ThreadListCountOfBoard(infra.S.Conn, board_id)
	thread_list := threads.BoardThreadPreviewList(infra.S.Conn, board_id, page_limit, offset)
	pageCount, pageHelper := infra.GeneratePagination(total, page_limit)

	data := map[string]any{
		"thread_list": thread_list,
		"page_count":  pageCount,
		"page_helper": pageHelper,
		"page":        page,
		"base_url":    fmt.Sprintf("/admin/boards/%d/threads/", board_id),
	}

	infra.S.HTML(w, r, http.StatusOK, "back/board_list_thread_list", data)
}

func BoardOneAdd(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.BoardOneAdd"

	data := map[string]any{
		"board":      boards.Board{},
		"topic_list": topics.TopicList(infra.S.Conn),
	}

	infra.S.HTML(w, r, http.StatusOK, "back/board_one", data)
}

func BoardOne(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.BoardOne"
	redirect_url := "/admin/boards/"
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil || id < 1 {
		logs.LogWarn(op, errors.Str("No id provided!"))
		infra.S.Redirect(w, r, redirect_url)
		return
	}

	data := map[string]any{
		"board":      boards.BoardOne(infra.S.Conn, id),
		"topic_list": topics.TopicList(infra.S.Conn),
	}

	infra.S.HTML(w, r, http.StatusOK, "back/board_one", data)
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
	bd := boards.Board{
		Name:        r.PostFormValue("name"),
		Topic_id:    topic_id,
		Code:        r.PostFormValue("code"),
		Description: r.PostFormValue("description"),
		Deleted_at:  deleted_at,
	}

	val, ok := vars["id"]
	if ok { // update
		id, err := strconv.Atoi(val)
		if err != nil || id < 1 {
			logs.LogWarn(op, errors.Str("No id provided!"))
			infra.S.Redirect(w, r, "/admin/boards/")
			return
		}
		bd.ID = id
	}

	boards.BoardOneUpdate(infra.S.Conn, &bd)

	redirect_url := fmt.Sprintf("/admin/boards/%d/", bd.ID)
	infra.S.Redirect(w, r, redirect_url)
}

func BoardOneCreate(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.BoardOneCreate"
	redirect_url := "/admin/boards/"

	topic_id, _ := strconv.Atoi(r.PostFormValue("topic_id"))
	newBoard := boards.Board{
		Name:        r.PostFormValue("name"),
		Topic_id:    topic_id,
		Code:        r.PostFormValue("code"),
		Description: r.PostFormValue("description"),
		Deleted_at:  0,
	}

	boards.BoardOneCreate(infra.S.Conn, &newBoard)

	infra.S.Redirect(w, r, redirect_url)
}
