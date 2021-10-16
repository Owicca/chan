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

func init() {
	adminRouter := infra.S.Router.PathPrefix("/admin").Subrouter()
	adminRouter.HandleFunc("/users/", http.HandlerFunc(UserList)).Methods(http.MethodGet).Name("user_list")
	adminRouter.HandleFunc("/users/{user_id:[0-9]+}/", http.HandlerFunc(UserOne)).Methods(http.MethodGet).Name("user_one")
	adminRouter.HandleFunc("/users/{user_id:[0-9]+}/", http.HandlerFunc(UserOneUpdate)).Methods(http.MethodPost).Name("user_one_update")
}

func UserList(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.UserList"
	data := map[string]interface{} {
		"users": users.UserList(infra.S.Conn),
	}

	infra.S.HTML(w, http.StatusOK, "back/user_list", data)
}

func UserOne(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.UserOne"
	vars := mux.Vars(r)

	user_id, err := strconv.Atoi(vars["user_id"])
	if err != nil || user_id < 1 {
		logs.LogWarn(op, errors.Str("No user_id provided!"))
		infra.S.HTML(w, http.StatusOK, "back/user", nil)
		return
	}

	data := map[string]interface{} {
		"user": users.UserOne(infra.S.Conn, user_id),
		"roles": acl.RoleList(infra.S.Conn),
		"statusList": users.UserStatusList(),
	}

	infra.S.HTML(w, http.StatusOK, "back/user", data)
}

func UserOneUpdate(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.UserOneUpdate"
	if err := r.ParseForm(); err != nil {
		logs.LogErr(op, err)
		infra.S.HTML(w, http.StatusOK, "back/user", nil)
		return
	}

	user_id, err := strconv.Atoi(r.PostFormValue("user_id"))
	if err != nil || user_id < 1 {
		logs.LogWarn(op, errors.Str("No user_id provided!"))
		infra.S.HTML(w, http.StatusOK, "back/user", nil)
		return
	}

	role_id, _ := strconv.Atoi(r.PostFormValue("role"))

	users.UserOneUpdate(infra.S.Conn, users.User{
		ID: user_id,
		Username: r.PostFormValue("username"),
		Email: r.PostFormValue("email"),
		Status: r.PostFormValue("status"),
		RoleId: role_id,
	})

	data := map[string]interface{} {
		"user": users.UserOne(infra.S.Conn, user_id),
		"roles": acl.RoleList(infra.S.Conn),
		"statusList": users.UserStatusList(),
	}

	infra.S.HTML(w, http.StatusOK, "back/user", data)
}