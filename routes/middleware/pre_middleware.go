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
	LoadPreMd(infra.S)
}

// Load middlewares
func LoadPreMd(srv *infra.Server) {
	srv.Router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		LogRequest(w, r)
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
			infra.S.Data["request"] = r
			next.ServeHTTP(w, r)
		})
	})

	srv.Router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, _ := infra.S.SessionStore.Get(r, infra.S.Config.Sessions.Key)
			//log.Println("get session pre", session.Values["user"])
			infra.S.Data = infra.MergeMapsInterface(infra.S.Data, session.Values)
			next.ServeHTTP(w, r)
		})
	})

	srv.Router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//if _, ok := srv.Data["user"]; !ok {
			//	urlPath := r.URL.Path
			//	for _, url := range sessions.PublicUrl {
			//		if url == urlPath {
			//			next.ServeHTTP(w, r)
			//			return
			//		}
			//	}
			//	infra.S.Redirect(w, r, "/admin/login/")
			//	return
			//}

			//user_id := srv.Data["user"].(map[string]any)["ID"]

			//session, _ := srv.SessionStore.Get(r, strconv.Itoa(user_id.(int)))
			//flashList := session.Flashes()
			//if len(flashList) > 0 {
			//	srv.Data["flash_list"] = flashList
			//}
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
