package backend

import (
	"github.com/Owicca/chan/models/users"
	"github.com/Owicca/chan/models/acl"
	"github.com/Owicca/chan/models/logs"
	"github.com/Owicca/chan/infra"
	"net/http"
	"github.com/gorilla/mux"

	"strconv"
	"upspin.io/errors"
)

type viewUser struct {
	ID int
	Username string
	Email string
	Role string
}

func UserList(srv *infra.Server) http.HandlerFunc {
	const op errors.Op = "back.UserOne"

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{} {
			"users": users.UserList(srv.Conn),
		}

		srv.Render(w, http.StatusOK, "back/user_list", data)
	})
}

func UserOne(srv *infra.Server) http.HandlerFunc {
	const op errors.Op = "back.UserOne"

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		user_id, err := strconv.Atoi(vars["user_id"])
		if err != nil || user_id < 1 {
			logs.LogWarn(op, errors.Str("No user_id provided!"))
			srv.Render(w, http.StatusOK, "back/user", nil)
			return
		}

		data := map[string]interface{} {
			"user": users.UserOne(srv.Conn, user_id),
			"roles": acl.RoleList(srv.Conn),
			"statusList": users.UserStatusList(),
		}

		srv.Render(w, http.StatusOK, "back/user", data)
	})
}

func UserOneUpdate(srv *infra.Server) http.HandlerFunc {
	const op errors.Op = "back.UserOneUpdate"

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user_id, err := strconv.Atoi(r.PostFormValue("user_id"))
		if err != nil {
			logs.LogWarn(op, errors.Str("No user_id provided!"))
			srv.Render(w, http.StatusOK, "back/user", nil)
			return
		}

		if err := r.ParseForm(); err != nil {
			logs.LogErr(op, err)
			srv.Render(w, http.StatusOK, "back/user", nil)
			return
		}

		data := map[string]interface{} {
			"user": users.UserOne(srv.Conn, user_id),
			"roles": acl.RoleList(srv.Conn),
			"statusList": users.UserStatusList(),
		}

		srv.Render(w, http.StatusOK, "back/user", data)
	})
}