package backend

import (
	"net/http"
	"upspin.io/errors"

	"github.com/Owicca/chan/models/boards"

	"github.com/Owicca/chan/infra"
)

func init() {
	adminRouter := infra.S.Router.PathPrefix("/admin").Subrouter()
	adminRouter.HandleFunc("/", http.HandlerFunc(Index)).Methods(http.MethodGet).Name("back_index")
}

func Index(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.Index"
	data := map[string]interface{} {
		"title": "The title",
		"topics": map[string]interface{} {
			"Topic1": boards.BoardList(infra.S.Conn),
		},

	}

	infra.S.HTML(w, http.StatusOK, "back/index", data)
}