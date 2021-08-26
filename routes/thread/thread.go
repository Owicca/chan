package thread

import (
	"net/http"

	"github.com/Owicca/chan/infra"
)


func Thread(srv *infra.Server) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{} {
			"posts": []struct{
				Name string
				Content string
			} {
				{
					Name: "p1",
					Content: "p1",
				},
				{
					Name: "p2",
					Content: "p2",
				},
			},
		}

		srv.Render(w, http.StatusOK, "front/thread_index", data)
	})
}