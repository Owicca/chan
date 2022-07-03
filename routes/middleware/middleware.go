package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Owicca/chan/infra"
	"github.com/Owicca/chan/models/sessions"

	"go.uber.org/zap"
)

func init() {
	LoadMd(infra.S)
}

// Load middlewares
func LoadMd(srv *infra.Server) {
	srv.Router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		LogRequest(w, r)
		template404Path := "front/404"
		if strings.HasPrefix(r.URL.Path, "/admin") {
			template404Path = "back/404"
		}
		srv.HTML(w, r, http.StatusNotFound, template404Path, nil)
		return
	})
	srv.Router.Use(setCSPHeader)

	// assets
	// web server will handle assets from now on
	//srv.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))).Methods(http.MethodGet).Name("static")

	srv.Router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			LogRequest(w, r)
			infra.S.Data["page_limit"] = 15
			next.ServeHTTP(w, r)
		})
	})

	srv.Router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !strings.HasPrefix(r.URL.Path, "/admin") {
				next.ServeHTTP(w, r)
				return
			}

			session, _ := srv.SessionStore.Get(r, srv.Config.Sessions.Key)
			//log.Printf("read %+v\n", session)
			if session.IsNew {
				urlPath := r.URL.Path
				for _, url := range sessions.PublicUrl {
					if url == urlPath {
						//log.Println("is public")
						next.ServeHTTP(w, r)
						return
					}
				}
				//log.Println("not logged")
				infra.S.Redirect(w, r, "/admin/login/")
				return
			}

			srv.Session = session
			user_id, _ := session.Values["user_id"].(int)
			ses := sessions.Get(srv.Conn, user_id)

			if err := json.Unmarshal([]byte(ses.Data), &srv.Data); err != nil {
				log.Fatalf("err while unmarshaling db session data (%s)", err)
			}

			next.ServeHTTP(w, r)
		})
	})

	srv.Router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if srv.Session != nil {
				flashList := srv.Session.Flashes()
				if len(flashList) > 0 {
					srv.Data["flash_list"] = flashList
				}
			}
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
