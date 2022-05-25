package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Owicca/chan/infra"

	"go.uber.org/zap"
)

func init() {
	LoadMd(infra.S)
}

// Load middlewares
func LoadMd(srv *infra.Server) {
	srv.Router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		LogRequest(w, r)
		// res := map[string]any{
		// 	"success": false,
		// 	"data":    nil,
		// 	"message": "Not found!",
		// }
		// srv.JSON(w, http.StatusNotFound, res)
		template404Path := "front/404"
		if strings.HasPrefix(r.URL.Path, "/admin") {
			template404Path = "back/404"
		}
		srv.HTML(w, http.StatusNotFound, template404Path, nil)
		return
	})
	srv.Router.Use(setCSPHeader)

	// assets
	// web server will handle assets from now on
	//srv.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))).Methods(http.MethodGet).Name("static")

	srv.Router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			LogRequest(w, r)
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

func LogRequest(w http.ResponseWriter, r *http.Request) {
	timestamp := time.Now().Unix()
	url := r.RequestURI
	remote_addr := r.RemoteAddr
	method := r.Method

	logMsg := fmt.Sprintf("%s %s %s", remote_addr, method, url)
	zap.L().Info(logMsg, zap.Int64("timestamp", timestamp))
}
