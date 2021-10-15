package routes

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Owicca/chan/infra"
	"github.com/Owicca/chan/routes/index"
	"github.com/Owicca/chan/routes/board"
	"github.com/Owicca/chan/routes/thread"

	b "github.com/Owicca/chan/routes/backend"

	"go.uber.org/zap"
	// "log"
)

// Load routes and middlewares
func LoadRoutes(srv *infra.Server) {
	srv.Router.NotFoundHandler = NotFound(srv)
	srv.Router.Use(setCSPHeader)

	// assets
	srv.RegisterPathPrefix("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))), []string{http.MethodGet}, "static")

	// frontend
	srv.RegisterRoute("/", index.Index(srv), []string{http.MethodGet}, "site_index")
	srv.RegisterRoute("/{board_name:[a-z]+}/", board.Board(srv), []string{http.MethodGet}, "board_index")
	srv.RegisterRoute("/{board_name:[a-z]+}/thread/{thread_id:[0-9]+}/", thread.Thread(srv), []string{http.MethodGet}, "thread_index")

	// backend
	adminRouter := srv.Router.PathPrefix("/admin").Subrouter()
	srv.RegisterSubRoute(adminRouter, "/users/", b.UserList(srv), []string{http.MethodGet}, "user_list")
	srv.RegisterSubRoute(adminRouter, "/users/{user_id:[0-9]+}/", b.UserOne(srv), []string{http.MethodGet}, "user_one")
	srv.RegisterSubRoute(adminRouter, "/users/{user_id:[0-9]+}/", b.UserOneUpdate(srv), []string{http.MethodPost}, "user_one_update")

	srv.Router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			LogRequest(srv, w, r)
			next.ServeHTTP(w, r)
		})
	})
}

func setCSPHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		csp := []string{
			"default-src *",
			//"script-src 'self'",
			//"connect-src 'self'",
			//"img-src 'self'",
			//"style-src 'self'",
			//"base-uri 'self'",
			//"form-action 'self'",
		}

		header.Set("Content-Security-Policy", strings.Join(csp, ";"))
		// header.Set("Access-Control-Allow-Origin", "*")

		next.ServeHTTP(w, r)
	})
}

// Created this just so the 
func LogRequest(srv *infra.Server, w http.ResponseWriter, r *http.Request) {
	timestamp := time.Now().Unix()
	url := r.RequestURI
	remote_addr := r.RemoteAddr
	method := r.Method

	logMsg := fmt.Sprintf("%s %s %s", remote_addr, method, url)
	zap.L().Info(logMsg, zap.Int64("timestamp", timestamp))
}

// 404 handler
func NotFound(srv *infra.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		LogRequest(srv, w, r)
		// res := map[string]interface{}{
		// 	"success": false,
		// 	"data":    nil,
		// 	"message": "Not found!",
		// }
		// srv.JSON(w, http.StatusNotFound, res)
		srv.Render(w, http.StatusNotFound, "back/404", nil)
		return
	}
}
