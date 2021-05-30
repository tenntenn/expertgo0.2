package list9

import (
	"io"
	"unicode"
	"unicode/utf8"
)

func UpperCount(r io.Reader) (int, error) {
	var count int
	b, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}

	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		b = b[size:]
		if r != utf8.RuneError && unicode.IsUpper(r) {
			count++
		}
	}

	return count, nil
}
