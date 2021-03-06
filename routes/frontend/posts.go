package frontend

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Owicca/chan/infra"
	"github.com/Owicca/chan/models/logs"
	"github.com/Owicca/chan/models/media"
	"github.com/Owicca/chan/models/posts"
	"github.com/gorilla/mux"
	"upspin.io/errors"
)

func init() {
	infra.S.HandleFunc("/boards/{board_code:[a-z0-9]+}/threads/{thread_id:[0-9]+}/", PostList).Methods(http.MethodGet).Name("f_post_list")
	infra.S.HandleFunc("/boards/{board_code:[a-z0-9]+}/threads/{thread_id:[0-9]+}/{page:[0-9]+}/", PostList).Methods(http.MethodGet).Name("f_post_list")
	infra.S.HandleFunc("/boards/{board_code:[a-z0-9]+}/threads/{thread_id:[0-9]+}/", CreatePost).Methods(http.MethodPost).Name("f_create_post")
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "front.CreatePost"
	vars := mux.Vars(r)
	redirect_url := fmt.Sprintf("/boards/%s/threads/%s/", vars["board_code"], vars["thread_id"])
	if err := r.ParseMultipartForm(media.MaxFileSize); err != nil {
		logs.LogErr(op, err)

		infra.S.Errors.Set("misc", "Invalid file size!")
		infra.S.Redirect(w, r, redirect_url)
		return
	}

	thread_id, err := strconv.Atoi(vars["thread_id"])
	if err != nil {
		logs.LogErr(op, err)

		infra.S.Errors.Set("misc", "Invalid thread id!")
		infra.S.Redirect(w, r, redirect_url)
		return
	}
	newPost := posts.Post{
		Created_at: time.Now().Unix(),
		Status:     string(posts.PostStatusActive),
		Thread_id:  thread_id,
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
		infra.S.Errors.Set("content", "No content provided!")
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
			logs.LogWarn(op, errors.Str("Invalid name provided!"))
			infra.S.Errors.Set("name", "Invalid name format!")
			infra.S.Redirect(w, r, redirect_url)
			return
		}
	}

	posts.PostOneCreate(infra.S.Conn, &newPost)

	mediaList, ok := r.MultipartForm.File["media"]
	if ok {
		m := mediaList[0]
		mediaFile, err := m.Open()
		if err != nil {
			logs.LogErr(op, err)

			infra.S.Errors.Set("media", "Error while processing the file!")
			posts.PostOneDelete(infra.S.Conn, newPost.ID)
			infra.S.Redirect(w, r, redirect_url)
			return
		}
		defer mediaFile.Close()
		if m.Size < media.MinFileSize && m.Size > media.MaxFileSize {
			errMsg := fmt.Sprintf("File size should be between %d KB and %d!", media.MinFileSize*1000, media.MaxFileSize*1000*1000)

			logs.LogErr(op, errMsg)
			infra.S.Errors.Set("media", errMsg)
			posts.PostOneDelete(infra.S.Conn, newPost.ID)
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
			infra.S.Errors.Set("media", "Error while uploading the file!")
			posts.PostOneDelete(infra.S.Conn, newPost.ID)
		} else {
			infra.S.Conn.Create(&newMedia)
		}
	}

	infra.S.Redirect(w, r, redirect_url)
}

func PostList(w http.ResponseWriter, r *http.Request) {
	const op errors.Op = "front.PostList"
	vars := mux.Vars(r)

	thread_id, err := strconv.Atoi(vars["thread_id"])
	if err != nil {
		logs.LogWarn(op, errors.Str("No thread_id provided!"))
		infra.S.Redirect(w, r, infra.S.GenerateUrl(fmt.Sprintf("/boards/%s/threads/", vars["board_code"])))
		return
	}
	post_list := posts.ThreadPostList(infra.S.Conn, thread_id, 0, 0)
	reply_count := len(post_list) - 1
	if reply_count < 0 {
		reply_count = 0
	}

	media_count := 0
	for _, post := range post_list {
		if post.Media.Object_id > 0 {
			media_count += 1
		}
	}

	data := map[string]any{
		"post_list": post_list,
		"stats": map[string]any{
			"reply_count": reply_count,
			"media_count": media_count,
		},
		"board_code": vars["board_code"],
		"thread_id":  vars["thread_id"],
	}

	infra.S.HTML(w, r, http.StatusOK, "front/post_list", data)
}
