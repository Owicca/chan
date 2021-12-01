package backend

import (
	"github.com/Owicca/chan/infra"
	"github.com/Owicca/chan/models/boards"
	"net/http"

	"upspin.io/errors"

	"github.com/Owicca/chan/models/logs"
	"strconv"
	"github.com/gorilla/mux"
)

func init() {
	adminRouter := infra.S.Router.PathPrefix("/admin").Subrouter()
	adminRouter.HandleFunc("/boards/", http.HandlerFunc(BoardList)).Methods(http.MethodGet).Name("board_list")
	adminRouter.HandleFunc("/boards/{board_id:[0-9]+}/", http.HandlerFunc(BoardOne)).Methods(http.MethodGet).Name("board_one")
	adminRouter.HandleFunc("/boards/{board_id:[0-9]+}/", http.HandlerFunc(BoardOneUpdate)).Methods(http.MethodPost).Name("board_one_update")
}

func BoardList(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.BoardList"

	data := map[string]interface{} {
		"boards": boards.BoardListWithThreadCount(infra.S.Conn),
	}

	infra.S.HTML(w, http.StatusOK, "back/board_list", data)
}

func BoardOne(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.BoardOne"
	vars := mux.Vars(r)

	board_id, err := strconv.Atoi(vars["board_id"])
	if err != nil || board_id < 1 {
		logs.LogWarn(op, errors.Str("No board_id provided!"))
		infra.S.HTML(w, http.StatusOK, "back/board_one", nil)
		return
	}

	data := map[string]interface{} {
		"board": boards.BoardOne(infra.S.Conn, board_id),
	}

	infra.S.HTML(w, http.StatusOK, "back/board_one", data)
}

func BoardOneUpdate(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.BoardOneUpdate"
	if err := r.ParseForm(); err != nil {
		logs.LogErr(op, err)
		infra.S.HTML(w, http.StatusOK, "back/board_one", nil)
		return
	}

	board_id, err := strconv.Atoi(r.PostFormValue("board_id"))
	if err != nil || board_id < 1 {
		logs.LogWarn(op, errors.Str("No board_id provided!"))
		infra.S.HTML(w, http.StatusOK, "back/board_one", nil)
		return
	}

	boards.BoardOneUpdate(infra.S.Conn, boards.Board{
		ID: board_id,
		Name: r.PostFormValue("name"),
		Code: r.PostFormValue("code"),
		Description: r.PostFormValue("description"),
	})

	data := map[string]interface{} {
		"board": boards.BoardOne(infra.S.Conn, board_id),
	}

	infra.S.HTML(w, http.StatusOK, "back/board_one", data)
}