package acl

import(
	"testing"
)

func CheckTest(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		t.Errorf("%s", "error")
	})
	t.Run("empty values", func(t *testing.T) {
		t.Errorf("%s", "error")
	})
	t.Run("invalid values", func(t *testing.T) {
		t.Errorf("%s", "error")
	})
}