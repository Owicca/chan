package backend

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"

	"github.com/Owicca/chan/infra"
	"github.com/Owicca/chan/models/boards"
	"github.com/Owicca/chan/models/posts"
	"github.com/Owicca/chan/models/threads"
	"github.com/Owicca/chan/models/topics"
)

func init() {
	adminRouter := infra.S.Router.PathPrefix("/admin").Subrouter()
	adminRouter.HandleFunc("/settings/", SettingList).Methods(http.MethodGet).Name("setting_list")
	adminRouter.HandleFunc("/dummy/", LoadDummyData).Methods(http.MethodGet).Name("dummy")
}

func SettingList(w http.ResponseWriter, r *http.Request) {
	infra.S.HTML(w, r, http.StatusOK, "back/setting_list", nil)
}

func LoadDummyData(w http.ResponseWriter, r *http.Request) {
	infra.ClearDb(infra.S.Conn)
	pidReg := regexp.MustCompile(`#p(\d{1,})`)

	for id := 1; id <= 10; id++ {
		newTopic := topics.Topic{
			ID:   id,
			Name: fmt.Sprintf("topic_%d", id),
		}
		topics.TopicOneCreate(infra.S.Conn, &newTopic)
	}

	bList := infra.LoadBoards("./log/boards.json")
	bIdList := []int{}
	for _, b := range bList {
		if b.Ws_board == 0 {
			continue
		}

		newBoard := boards.Board{
			Name:        b.Title,
			Code:        b.Board,
			Description: b.Meta_description,
			Topic_id:    rand.Intn(10) + 1,
		}
		boards.BoardOneCreate(infra.S.Conn, &newBoard)
		bIdList = append(bIdList, newBoard.ID)
	}

	//pList := infra.LoadPosts("./log/posts.json")
	pCh := infra.LoadThreads("./log/threads.json")
	defer close(pCh)
	for p := range pCh {
		if p.No == 0 {
			break
		}
		log.Println(p.No)
		thread_id := p.Resto
		if p.Resto == 0 {
			boardIndex := rand.Intn(len(bIdList))

			thread_id = p.No
			newThread := threads.Thread{
				ID:       thread_id,
				Subject:  fmt.Sprintf("Thread subject: %s (thread_%d)", p.Name, thread_id),
				Board_id: bIdList[boardIndex],
			}
			threads.ThreadOneCreate(infra.S.Conn, &newThread)
		}

		var created_at int64 = int64(p.Time)
		newPost := posts.Post{
			ID:         p.No,
			Created_at: created_at,
			Status:     string(posts.PostStatusActive),
			Thread_id:  thread_id,
			Name:       infra.P.Sanitize(p.Name),
			Content:    infra.P.Sanitize(p.Com),
		}
		posts.PostOneCreate(infra.S.Conn, &newPost)

		matches := pidReg.FindAllStringSubmatch(p.Com, -1)
		for _, m := range matches {
			if len(m) > 1 {
				dest, _ := strconv.Atoi(m[1])
				link := posts.Link{
					Src:  p.No,
					Dest: dest,
				}

				posts.LinkOneCreate(infra.S.Conn, &link)
			}
		}
	}

	log.Println("finished entering dummy data")

	infra.S.Redirect(w, r, "/admin/settings/")
}
