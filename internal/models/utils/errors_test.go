package utils

import (
	"reflect"
	"testing"
)

func TestErrorsGetSet(t *testing.T) {
	data := []struct {
		k string
		v []any
	}{
		{"k1", []any{"str1"}},
		{"k2", []any{2}},
		{"k3", []any{NewErrors()}},
	}
	errors := NewErrors()

	for _, d := range data {
		errors.Set(d.k, d.v)

		got := errors.Get(d.k)
		expected := d.v

		if !reflect.DeepEqual(got, expected) {
			t.Fatalf("Did not get what was expected (%s, %s)", got, expected)
		}
	}
}
