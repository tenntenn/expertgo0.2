package list6

import (
	"bytes"
	"io"
)

func IsPNG(r io.Reader) (bool, error) {
	// PNG形式のマジックナンバー
	magicnum := []byte{137, 80, 78, 71}
	buf := make([]byte, len(magicnum))
	_, err := io.ReadFull(r, buf)
	if err != nil {
		return false, err
	}
	return bytes.Equal(magicnum, buf), nil
}
