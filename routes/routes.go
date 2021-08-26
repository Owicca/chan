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
	"go.uber.org/zap"
)

func LoadRoutes(srv *infra.Server) {
	// srv.Router = setRouterPrerequisites(srv.Router)
	srv.Router.NotFoundHandler = NotFound(srv)
	srv.Router.Use(setCSPHeader)

	srv.RegisterPathPrefix("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))), []string{http.MethodGet}, "static")
	srv.RegisterRoute("/", index.Index(srv), []string{http.MethodGet}, "site_index")
	srv.RegisterRoute("/{board_name:[a-z]+}/", board.Board(srv), []string{http.MethodGet}, "board_index")
	srv.RegisterRoute("/{board_name:[a-z]+}/thread/{thread_id:[0-9]+}/", thread.Thread(srv), []string{http.MethodGet}, "thread_index")
	RegisterRequestLogger(srv)
}

// func setRouterPrerequisites(r *mux.Router) *mux.Router {
// 	subrouter := r.PathPrefix("/api/v1").Headers("Content-Type", "application/json").Subrouter()

// 	return subrouter
// }

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
		header.Set("Access-Control-Allow-Origin", "*")

		next.ServeHTTP(w, r)
	})
}

func LogRequest(srv *infra.Server, w http.ResponseWriter, r *http.Request) {
	timestamp := time.Now().Unix()
	url := r.RequestURI
	remote_addr := r.RemoteAddr
	method := r.Method

	logMsg := fmt.Sprintf("%s %s %s", remote_addr, method, url)
	zap.L().Info(logMsg, zap.Int64("timestamp", timestamp))
}

func RegisterRequestLogger(srv *infra.Server) {
	srv.Router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			LogRequest(srv, w, r)
			next.ServeHTTP(w, r)
		})
	})
}

func NotFound(srv *infra.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		LogRequest(srv, w, r)
		// res := map[string]interface{}{
		// 	"success": false,
		// 	"data":    nil,
		// 	"message": "Not found!",
		// }
		// srv.JSON(w, http.StatusNotFound, res)
		srv.Render(w, http.StatusNotFound, "404.tpl", nil)
		return
	}
}
