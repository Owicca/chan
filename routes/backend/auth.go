package backend

import (
	"net/http"

	"github.com/Owicca/chan/infra"
	"github.com/Owicca/chan/models/logs"
	"github.com/Owicca/chan/models/users"
	"github.com/fatih/structs"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"upspin.io/errors"
)

var (
	adminRouter   *mux.Router
	SessionMaxAge = 86400 * 7
)

func init() {
	adminRouter = infra.S.PathPrefix("/admin/").Subrouter()
	adminRouter.HandleFunc("/login/", LoginForm).Methods(http.MethodGet).Name("login_form")
	adminRouter.HandleFunc("/login/", Login).Methods(http.MethodPost).Name("login")
	adminRouter.HandleFunc("/logout/", Logout).Methods(http.MethodPost).Name("logout")
}

func LoginForm(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.Login"
	infra.S.HTML(w, http.StatusOK, "back/login", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "back.Login"
	url, _ := adminRouter.Get("login_form").URL()

	email, pass1, pass2 := r.PostFormValue("email"), r.PostFormValue("password1"), r.PostFormValue("password2")
	if err := users.UserValidate(email, pass1, pass2); err != nil {
		logs.LogErr(op, err)

		infra.S.Redirect(w, r, url.String())
		return
	}

	user, err := users.UserGetByCredentials(infra.S.Conn, email, pass1)
	if err != nil {
		logs.LogErr(op, err)

		infra.S.Redirect(w, r, url.String())
		return
	}

	session, _ := infra.S.SessionStore.Get(r, infra.S.Config.Sessions.Key)
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   SessionMaxAge,
		HttpOnly: true,
	}
	infra.S.Data["user"] = structs.Map(user)

	infra.S.Redirect(w, r, "/admin/")
}

func Logout(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "auth.Logout"
	session, _ := infra.S.SessionStore.Get(r, infra.S.Config.Sessions.Key)

	delete(session.Values, "user")
	if err := session.Save(r, w); err != nil {
		logs.LogErr(op, errors.Errorf("Could not save session on logout (%s)!", err))
	}

	infra.S.Redirect(w, r, "/admin/login/")
}