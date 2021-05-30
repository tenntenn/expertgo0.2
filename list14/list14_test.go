package list14

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestUpperCount(t *testing.T) {
	str, want := "AbcD", 2
	var buf bytes.Buffer
	r := io.TeeReader(strings.NewReader(str), &buf)
	got, err := UpperCount(r)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Error(want, "!=", got)
	}
}
