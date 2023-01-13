package frontend

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Owicca/chan/infra"
	"github.com/Owicca/chan/models/boards"
	"github.com/Owicca/chan/models/logs"
	"github.com/Owicca/chan/models/media"
	"github.com/Owicca/chan/models/posts"
	"github.com/Owicca/chan/models/threads"

	"github.com/gorilla/mux"
	"upspin.io/errors"
)

func init() {
	infra.S.HandleFunc("/boards/{board_code:[a-z0-9]+}/", ThreadList).Methods(http.MethodGet).Name("board_thread_list")
	infra.S.HandleFunc("/boards/{board_code:[a-z0-9]+}/{page:[0-9]+}/", ThreadList).Methods(http.MethodGet).Name("board_thread_list_page")
	infra.S.HandleFunc("/boards/{board_code:[a-z0-9]+}/", ThreadCreate).Methods(http.MethodPost).Name("f_board_create")
}

func ThreadCreate(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "front.ThreadCreate"
	vars := mux.Vars(r)
	redirect_url := fmt.Sprintf("/boards/%s/", vars["board_code"])
	if err := r.ParseMultipartForm(media.MaxFileSize); err != nil {
		logs.LogErr(op, errors.Errorf("Could not parse form (%s)", err))
		infra.S.Errors.Set("misc", []any{errors.Str("Invalid file size!")})
		infra.S.Redirect(w, r, redirect_url)
		return
	}

	subject := r.PostFormValue("subject")
	if subject == "" {
		err := errors.Str("A subject is required when creating a thread!")
		logs.LogWarn(op, err)
		infra.S.Errors.Set("subject", []any{err})
		infra.S.Redirect(w, r, redirect_url)
		return
	}
	mediaList, ok := r.MultipartForm.File["media"]
	if !ok || len(mediaList) == 0 {
		err := errors.Errorf("Media is required when creating a thread!")
		logs.LogErr(op, err)
		infra.S.Errors.Set("media", []any{err})
		infra.S.Redirect(w, r, redirect_url)
		return
	}
	board_id := boards.BoardIdByCode(infra.S.Conn, vars["board_code"])
	thread := threads.Thread{
		Board_id: board_id,
		Subject:  subject,
	}
	threads.ThreadOneCreate(infra.S.Conn, &thread)
	newPost := posts.Post{
		Created_at: time.Now().Unix(),
		Status:     string(posts.PostStatusActive),
		Thread_id:  thread.ID,
	}

	var (
		trip    string
		secure  string
		name    string
		inp     string
		content string
	)

	content = r.PostFormValue("content")
	if content == "" {
		logs.LogWarn(op, errors.Str("No content provided!"))
		infra.S.Errors.Set("content", []any{errors.Str("No content provided!")})
		infra.S.Redirect(w, r, redirect_url)
		return
	}
	newPost.Content = content

	inp = r.PostFormValue("name")
	if inp != "" {
		name, trip, secure = posts.DeconstructInput(inp)
		if name != "" && secure != "" {
			newPost.Name = name
			newPost.SecureTripcode = secure
		} else if name != "" && trip != "" {
			newPost.Name = name
			newPost.Tripcode = trip
		} else {
			err := errors.Str("Invalid name provided!")
			logs.LogWarn(op, err)
			infra.S.Errors.Set("name", []any{err})
			infra.S.Redirect(w, r, redirect_url)
			return
		}
	}

	posts.PostOneCreate(infra.S.Conn, &newPost)

	thread.Primary_post_id = newPost.ID
	threads.ThreadOneUpdate(infra.S.Conn, thread)

	m := mediaList[0]
	mediaFile, err := m.Open()
	if err != nil {
		logs.LogErr(op, err)

		infra.S.Errors.Set("media", []any{errors.Str("Error while processing the file!")})
		infra.S.Redirect(w, r, redirect_url)
		return
	}
	newMedia, err := media.CreateMedia(&media.Media{
		Object_id:   newPost.ID,
		Object_type: media.PostsObject,
		Name:        m.Filename,
		Size:        m.Size,
	}, mediaFile)
	if err != nil {
		logs.LogErr(op, err)
		infra.S.Errors.Set("media", []any{errors.Str("Error while creating the file!")})
	} else {
		infra.S.Conn.Create(&newMedia)
	}

	infra.S.Redirect(w, r, redirect_url)
}

func ThreadList(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "front.ThreadList"
	vars := mux.Vars(r)

	offset, _ := strconv.Atoi(vars["page"])

	totalThreads := threads.TotalActiveThreads(infra.S.Conn, vars["board_code"])

	pageCount, pageHelper := infra.GeneratePagination(totalThreads, threads.ThreadPageLimit)

	data := map[string]any{
		"thread_list": threads.ThreadPreviewByCode(infra.S.Conn, vars["board_code"], threads.ThreadPageLimit, offset),
		"board_code":  vars["board_code"],
		"page_count":  pageCount,
		"page_helper": pageHelper,
		"page":        offset,
	}

	infra.S.HTML(w, r, http.StatusOK, "front/thread_list", data)
}
