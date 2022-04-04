package backend

import (
	"net/http"
	"strconv"
	"time"

	"upspin.io/errors"
	"github.com/gorilla/mux"

	"github.com/Owicca/chan/models/logs"
	"github.com/Owicca/chan/infra"
	"github.com/Owicca/chan/models/topics"
)

func init() {
	adminRouter := infra.S.Router.PathPrefix("/admin").Subrouter()
	adminRouter.HandleFunc("/topics/", http.HandlerFunc(TopicList)).Methods(http.MethodGet).Name("topic_list")

	adminRouter.HandleFunc("/topics/add/", http.HandlerFunc(TopicOneAdd)).Methods(http.MethodGet).Name("topic_one_add")
	adminRouter.HandleFunc("/topics/", http.HandlerFunc(TopicOneUpdate)).Methods(http.MethodPost).Name("topic_one_add_post")

	adminRouter.HandleFunc("/topics/{id:[0-9]+}/", http.HandlerFunc(TopicOne)).Methods(http.MethodGet).Name("topic_one")
	adminRouter.HandleFunc("/topics/{id:[0-9]+}/", http.HandlerFunc(TopicOneUpdate)).Methods(http.MethodPost).Name("topic_one_post")
}

func TopicList(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back/TopicList"

	data := map[string]any {
		"topic_list": topics.TopicList(infra.S.Conn),
	}

	infra.S.HTML(w, http.StatusOK, "back/topic_list", data)
}

func TopicOneAdd(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back/TopicOneAdd"

	data := map[string]any {
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

	data := map[string]any {
		"topic": topics.TopicOne(infra.S.Conn, id),
	}

	infra.S.HTML(w, http.StatusOK, "back/topic_one", data)
}

func TopicOneUpdate(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back/TopicOneUpdate"
	vars := mux.Vars(r)

	form_deleted_at := r.PostFormValue("deleted_at")
	var deleted_at int64 = 0
	if form_deleted_at == "1" {
		deleted_at = time.Now().Unix()
	}

	val, ok := vars["id"]
	the_id := 0
	if ok {// update
		id, err := strconv.Atoi(val)
		if err != nil || id < 1 {
			logs.LogWarn(op, errors.Str("No id provided!"))
			infra.S.HTML(w, http.StatusOK, "back/topic_one_add", nil)
			return
		}
		the_id = id

		topics.TopicOneUpdate(infra.S.Conn, topics.Topic{
			ID: id,
			Name: r.PostFormValue("name"),
			Deleted_at: deleted_at,
		})
	} else {// create
		topics.TopicOneCreate(infra.S.Conn, topics.Topic{
			Name: r.PostFormValue("name"),
			Deleted_at: 0,
		})
	}

	data := map[string]any {
		"topic": topics.TopicOne(infra.S.Conn, the_id),
	}

	infra.S.HTML(w, http.StatusOK, "back/topic_one", data)
}