package backend

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/Owicca/chan/internal/infra"
	"github.com/Owicca/chan/internal/models/boards"
	"github.com/Owicca/chan/internal/models/posts"
	"github.com/Owicca/chan/internal/models/threads"
	"github.com/Owicca/chan/internal/models/topics"
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
	GenerateDummyData()

	infra.S.Redirect(w, r, "/admin/settings/")
}

func RandIntInRange(min int, max int) int {
	return rand.Intn(max-min) + min
}

func GenerateDummyData() {
	infra.ClearDb(infra.S.Conn)
	pidReg := regexp.MustCompile(`#p(\d{1,})`)

	for id := 1; id <= 10; id++ {
		lastYear := 1687807056
		now := int(time.Now().Unix())
		createdAt := RandIntInRange(lastYear, now)

		newTopic := topics.Topic{
			ID:         id,
			Created_at: int64(createdAt),
			Name:       fmt.Sprintf("topic_%d", id),
		}
		topics.TopicOneCreate(infra.S.Conn, &newTopic)
	}

	bList := infra.LoadBoards("./log/boards.json")
	bIdList := []int{}
	for _, b := range bList {
		if b.Ws_board == 0 {
			continue
		}

		lastYear := 1687807056
		now := int(time.Now().Unix())
		createdAt := RandIntInRange(lastYear, now)

		newBoard := boards.Board{
			Created_at:  int64(createdAt),
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

		lastYear := 1687807056
		now := int(time.Now().Unix())
		createdAt := RandIntInRange(lastYear, now)

		log.Println(p.No)
		thread_id := p.Resto
		if p.Resto == 0 {
			boardIndex := rand.Intn(len(bIdList))

			thread_id = p.No
			newThread := threads.Thread{
				ID:         thread_id,
				Created_at: int64(createdAt),
				Subject:    fmt.Sprintf("Thread subject: %s (thread_%d)", p.Name, thread_id),
				Board_id:   bIdList[boardIndex],
			}
			threads.ThreadOneCreate(infra.S.Conn, &newThread)
		}

		newPost := posts.Post{
			ID:         p.No,
			Created_at: int64(p.Time),
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
}
