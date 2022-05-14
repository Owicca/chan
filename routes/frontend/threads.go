package frontend

import (
	"net/http"

	"github.com/Owicca/chan/infra"
	"github.com/Owicca/chan/models/threads"

	"github.com/gorilla/mux"
	"upspin.io/errors"
)

func init() {
	infra.S.HandleFunc("/boards/{board_code:[a-z0-9]+}/", ThreadList).Methods(http.MethodGet).Name("board_thread_list")
}

func ThreadList(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "front.ThreadList"
	vars := mux.Vars(r)

	data := map[string]any{
		"thread_list": threads.BoardThreadListByCode(infra.S.Conn, vars["board_code"]),
		"board_code":  vars["board_code"],
	}

	infra.S.HTML(w, http.StatusOK, "front/thread_list", data)
}
