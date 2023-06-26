package backend

import (
	"net/http"

	"github.com/Owicca/chan/internal/infra"
)

func init() {
	adminRouter = infra.S.PathPrefix("/admin/").Subrouter()
	adminRouter.HandleFunc("/analytics/", Analytics).Methods(http.MethodGet).Name("analytics_list")
}

type Tab struct {
	ID       string
	Name     string
	IsActive bool
}

func Analytics(w http.ResponseWriter, r *http.Request) {

	data := map[string]any{
		"types": []Tab{
			Tab{"boards", "Boards", true},
			Tab{"links", "Links", false},
			Tab{"media", "Media", false},
			Tab{"objects", "Objects", false},
			Tab{"posts", "Posts", false},
			Tab{"threads", "Threads", false},
			Tab{"topics", "Topics", false},
			Tab{"users", "Users", false},
		},
	}

	infra.S.HTML(w, r, http.StatusOK, "back/analytics_list", data)
}
