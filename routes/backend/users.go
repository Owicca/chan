package backend

import (
	"github.com/Owicca/chan/models/users"
	"github.com/Owicca/chan/models/acl"
	"github.com/Owicca/chan/infra"
	"net/http"
	"github.com/gorilla/mux"

	// "log"
	"strconv"
)

type viewUser struct {
	ID int
	Username string
	Email string
	Role string
}

func UserList(srv *infra.Server) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{} {
			"users": users.GetUserList(srv.Conn),
		}

		srv.Render(w, http.StatusOK, "back/user_list", data)
	})
}

func UserOne(srv *infra.Server) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		user_id, err := strconv.Atoi(vars["user_id"])
		if err != nil {
			user_id = 0
		}

		data := map[string]interface{} {
			"user": users.GetUser(srv.Conn, user_id),
			"roles": acl.GetRoleList(srv.Conn),
			"statusList": users.UserStatusList(),
		}

		srv.Render(w, http.StatusOK, "back/user", data)
	})
}

func UserOneUpdate(srv *infra.Server) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		srv.Render(w, http.StatusOK, "back/user", nil)
	})
}