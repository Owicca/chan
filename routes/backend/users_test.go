package backend

import(
	"testing"
	"net/http/httptest"
	"io"
)

func TestUserList(t *testing.T) {
	method := "GET"
	url := "localhost:5900/admin/users/"
	
	req := httptest.NewRequest(method, url, nil)
	w := httptest.NewRecorder()
	UserList(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("%s UserList: could not read body content (%s)", method, err)
	}

	t.Log(resp.StatusCode)
	t.Log(string(body))
}