package middleware

import (
	"log"
	"net/http"

	"github.com/Owicca/chan/infra"
)

func init() {
	LoadPostMd(infra.S)
}

// Load middlewares
func LoadPostMd(srv *infra.Server) {
	srv.Router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, _ := infra.S.SessionStore.Get(r, infra.S.Config.Sessions.Key)
			//log.Println("get session post", session.Values)
			session.Values = infra.MergeMapsInterfaceReverse(session.Values, infra.S.Data)
			if err := session.Save(r, w); err != nil {
				log.Println("err", err)
			}
			next.ServeHTTP(w, r)
		})
	})
}
