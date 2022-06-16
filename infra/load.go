package infra

import (
	"encoding/json"
	"log"
	"os"
)

type JSCooldown struct {
	Threads int
	Replies int
	Images  int
}

type JSBoard struct {
	Board             string // code
	Title             string // name
	Ws_board          int
	Per_page          int
	Pages             int
	Max_filesize      int
	Max_comment_chars int
	Image_limit       int
	Cooldowns         JSCooldown
	Meta_description  string // description
	Is_archived       int
}

type JSPost struct {
	No           int // ID
	Resto        int // 0 if post is thread
	Sticky       int
	Closed       int
	Now          string // Created_at
	Name         string // Name
	Sub          string
	Com          string // Comment
	Filename     string
	Ext          string
	W            int
	H            int
	Tn_w         int
	Tn_h         int
	Tim          int
	Time         int
	Md5          string
	Fsize        int
	Capcode      string
	Semantic_url string
	Replies      int
	Images       int
	Unique_ips   int
}

func LoadBoards(path string) []JSBoard {
	fData, _ := os.ReadFile(path)

	data := struct {
		Boards []JSBoard
	}{}
	if err := json.Unmarshal(fData, &data); err != nil {
		log.Fatalf("err while unmarshaling (%s)", err)
	}

	return data.Boards
}

func LoadPosts(path string) []JSPost {
	fData, _ := os.ReadFile(path)

	data := struct {
		Posts []JSPost
	}{}
	if err := json.Unmarshal(fData, &data); err != nil {
		log.Fatalf("err while unmarshaling (%s)", err)
	}

	return data.Posts
}
