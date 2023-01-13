package backend

import (
	"fmt"
	"net/http"

	"github.com/Owicca/chan/infra"
	"github.com/Owicca/chan/models/acl"
	"github.com/Owicca/chan/models/boards"
	"github.com/Owicca/chan/models/logs"
	"github.com/Owicca/chan/models/sessions"
	"github.com/Owicca/chan/models/users"
	"github.com/gorilla/mux"

	"strconv"

	"upspin.io/errors"
)

func init() {
	adminRouter := infra.S.Router.PathPrefix("/admin").Subrouter()
	adminRouter.HandleFunc("/users/", UserList).Methods(http.MethodGet).Name("user_list")
	adminRouter.HandleFunc("/users/p{page:[0-9]+}/", UserList).Methods(http.MethodGet).Name("user_list_page")
	adminRouter.HandleFunc("/users/add/", UserCreateForm).Methods(http.MethodGet).Name("user_one_create")
	adminRouter.HandleFunc("/users/", UserOneCreate).Methods(http.MethodPost).Name("user_one_create")
	adminRouter.HandleFunc("/users/{user_id:[0-9]+}/", UserOne).Methods(http.MethodGet).Name("user_one")
	adminRouter.HandleFunc("/users/{user_id:[0-9]+}/", UserOneUpdate).Methods(http.MethodPost).Name("user_one_update")
}

func UserList(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.UserList"
	vars := mux.Vars(r)

	page_limit := int(infra.S.Data["page_limit"].(float64))
	page, _ := strconv.Atoi(vars["page"])
	offset := page * page_limit
	totalUsers := users.UserListCount(infra.S.Conn)
	users := users.UserList(infra.S.Conn, page_limit, offset)
	pageCount, pageHelper := infra.GeneratePagination(totalUsers, page_limit)

	data := map[string]any{
		"users":       users,
		"page_count":  pageCount,
		"page_helper": pageHelper,
		"page":        page,
		"base_url":    "/admin/users/",
	}

	infra.S.HTML(w, r, http.StatusOK, "back/user_list", data)
}

func UserOne(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.UserOne"
	vars := mux.Vars(r)

	user_id, err := strconv.Atoi(vars["user_id"])
	if err != nil || user_id < 1 {
		logs.LogWarn(op, errors.Str("No user_id provided!"))
		infra.S.HTML(w, r, http.StatusOK, "back/user", nil)
		return
	}

	data := map[string]any{
		"user":       users.UserOne(infra.S.Conn, user_id),
		"roles":      acl.RoleList(infra.S.Conn),
		"statusList": users.UserStatusList(),
	}

	infra.S.HTML(w, r, http.StatusOK, "back/user", data)
}

func UserCreateForm(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.UserCreateForm"

	data := map[string]interface{}{
		"user_status_list": users.UserStatusList(),
		"user_role_list":   acl.RoleList(infra.S.Conn),
		"board_list":       boards.BoardList(infra.S.Conn),
	}

	infra.S.HTML(w, r, http.StatusOK, "back/user_create_form", data)
}

func UserOneCreate(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.UserOneCreate"
	redirect_url := "/admin/users/add/"

	role_id_list := acl.RoleIdList(infra.S.Conn)
	role_id, err := strconv.Atoi(r.PostFormValue("role"))
	if err != nil || !infra.Contains(role_id_list, role_id) {
		logs.LogWarn(op, errors.Errorf("Invalid role id! (%s)", err))
		infra.S.Errors.Set("role", []any{errors.Str("Invalid role id!")})
		infra.S.Redirect(w, r, redirect_url)
		return
	}

	email := r.PostFormValue("email")
	status := r.PostFormValue("status")
	pass1 := r.PostFormValue("password1")
	pass2 := r.PostFormValue("password2")
	if err := users.UserValidate(email, pass1, pass2); err != nil {
		logs.LogWarn(op, errors.Errorf("Invalid email and pass! (%s)", err))
		infra.S.Errors.Set("password1", []any{err})
		infra.S.Redirect(w, r, redirect_url)
		return
	}

	pepper := sessions.GeneratePepper(32)
	hash := sessions.GeneratePassword(pass1, pepper)

	newUser := users.User{
		Username: r.PostFormValue("username"),
		Email:    email,
		Status:   status,
		Password: hash,
		Pepper:   pepper,
		RoleId:   role_id,
	}
	users.UserOneCreate(infra.S.Conn, &newUser)

	redirect_url = fmt.Sprintf("/admin/users/%d/", newUser.ID)
	infra.S.Redirect(w, r, redirect_url)
}

func UserOneUpdate(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.UserOneUpdate"
	default_redirect_url := "/admin/users/"

	user_id, err := strconv.Atoi(r.PostFormValue("user_id"))
	if err != nil || user_id < 1 {
		logs.LogWarn(op, errors.Str("No user_id provided!"))
		infra.S.Redirect(w, r, default_redirect_url)
		return
	}
	redirect_url := fmt.Sprintf("/admin/users/%d/", user_id)

	role_id, _ := strconv.Atoi(r.PostFormValue("role"))

	newUser := users.User{
		ID:       user_id,
		Username: r.PostFormValue("username"),
		Email:    r.PostFormValue("email"),
		Status:   r.PostFormValue("status"),
		RoleId:   role_id,
	}
	users.UserOneUpdate(infra.S.Conn, &newUser)

	infra.S.Redirect(w, r, redirect_url)
}
