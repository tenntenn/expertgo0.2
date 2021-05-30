package list6

import (
	"io"
	"testing"
)

// EOFを返さないReader
type neverEnding byte

func (b neverEnding) Read(p []byte) (n int, err error) {
	// bで埋める
	for i := range p {
		p[i] = byte(b)
	}
	return len(p), nil
}

func TestIsPNG(t *testing.T) {
	n, want := int64(10), false
	r := io.LimitReader(neverEnding('x'), n)
	got, err := IsPNG(r)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Error(want, "!=", got)
	}
}
