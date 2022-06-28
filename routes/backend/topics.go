package backend

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"upspin.io/errors"

	"github.com/Owicca/chan/infra"
	"github.com/Owicca/chan/models/logs"
	"github.com/Owicca/chan/models/topics"
)

func init() {
	adminRouter := infra.S.Router.PathPrefix("/admin").Subrouter()
	adminRouter.HandleFunc("/topics/", TopicList).Methods(http.MethodGet).Name("topic_list")
	adminRouter.HandleFunc("/topics/p{page:[0-9]+}/", TopicList).Methods(http.MethodGet).Name("topic_list_page")

	adminRouter.HandleFunc("/topics/add/", TopicOneAdd).Methods(http.MethodGet).Name("topic_one_add")
	adminRouter.HandleFunc("/topics/", TopicOneCreate).Methods(http.MethodPost).Name("topic_one_create")

	adminRouter.HandleFunc("/topics/{id:[0-9]+}/", TopicOne).Methods(http.MethodGet).Name("topic_one")
	adminRouter.HandleFunc("/topics/{id:[0-9]+}/", TopicOneUpdate).Methods(http.MethodPost).Name("topic_one_post")

	adminRouter.HandleFunc("/topics/{id:[0-9]+}/boards/", TopicOneBoardList).Methods(http.MethodGet).Name("topic_list_board_list")
}

func TopicList(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back/TopicList"
	vars := mux.Vars(r)

	page_limit := infra.S.Data["page_limit"].(int)
	page, _ := strconv.Atoi(vars["page"])
	offset := page * page_limit
	total := topics.TopicListCount(infra.S.Conn)
	topic_list := topics.TopicListWithBoardList(infra.S.Conn, page_limit, offset)
	pageCount, pageHelper := infra.GeneratePagination(total, page_limit)

	log.Println(total, len(topic_list))
	log.Println(page, offset)
	log.Println(pageCount, pageHelper)

	data := map[string]any{
		"topic_list":  topic_list,
		"page_count":  pageCount,
		"page_helper": pageHelper,
		"page":        page,
		"base_url":    "/admin/topics/",
	}

	infra.S.HTML(w, http.StatusOK, "back/topic_list", data)
}

func TopicOneAdd(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back/TopicOneAdd"

	data := map[string]any{
		"topic": topics.Topic{},
	}

	infra.S.HTML(w, http.StatusOK, "back/topic_one", data)
}

func TopicOne(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back/TopicOne"
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil || id < 1 {
		logs.LogWarn(op, errors.Str("No id provided!"))
		infra.S.HTML(w, http.StatusOK, "back/board_one", nil)
		return
	}

	data := map[string]any{
		"topic": topics.TopicOne(infra.S.Conn, id),
	}

	infra.S.HTML(w, http.StatusOK, "back/topic_one", data)
}

func TopicOneUpdate(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back/TopicOneUpdate"
	vars := mux.Vars(r)
	redirect_url := fmt.Sprintf("/admin/topics/%s/", vars["id"])

	form_deleted_at := r.PostFormValue("deleted_at")
	var deleted_at int64 = 0
	if form_deleted_at == "1" {
		deleted_at = time.Now().Unix()
	}
	tp := topics.Topic{
		Name:       r.PostFormValue("name"),
		Deleted_at: deleted_at,
	}

	val, ok := vars["id"]
	if ok { // update
		id, err := strconv.Atoi(val)
		if err != nil || id < 1 {
			logs.LogWarn(op, errors.Str("No id provided!"))
			infra.S.Redirect(w, r, "/admin/topics/")
			return
		}
		tp.ID = id
	}
	topics.TopicOneUpdate(infra.S.Conn, &tp)

	infra.S.Redirect(w, r, redirect_url)
}

func TopicOneCreate(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back/TopicOneCreate"
	redirect_url := "/admin/topics/"

	form_deleted_at := r.PostFormValue("deleted_at")
	var deleted_at int64 = 0
	if form_deleted_at == "1" {
		deleted_at = time.Now().Unix()
	}
	newTopic := topics.Topic{
		Name:       r.PostFormValue("name"),
		Deleted_at: deleted_at,
	}

	topics.TopicOneCreate(infra.S.Conn, &newTopic)

	infra.S.Redirect(w, r, redirect_url)
}

func TopicOneBoardList(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back/TopicOneCreate"
	vars := mux.Vars(r)

	topic_id := 0
	val, ok := vars["id"]
	if ok { // update
		id, err := strconv.Atoi(val)
		if err != nil || id < 1 {
			logs.LogWarn(op, errors.Str("No id provided!"))
			infra.S.Redirect(w, r, "/admin/topics/")
			return
		}
		topic_id = id
	}
	data := map[string]any{
		"topic": topics.TopicOneWithBoardList(infra.S.Conn, topic_id),
	}

	infra.S.HTML(w, http.StatusOK, "back/topic_one_board_list", data)
}
