package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func IsPNG(r io.ReadSeeker) (bool, error) {
	// PNG形式のマジックナンバー
	magicnum := []byte{137, 80, 78, 71}
	buf := make([]byte, len(magicnum))
	_, err := io.ReadFull(r, buf)
	if err != nil {
		return false, err
	}
	_, err = r.Seek(0, io.SeekStart)
	if err != nil {
		return false, err
	}
	return bytes.Equal(magicnum, buf), nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run() error {
	f, err := os.Open("gopher.png")
	if err != nil {
		return err
	}
	defer f.Close()
	isPNG, err := IsPNG(f)
	if err != nil {
		return err
	}
	if isPNG {
		fmt.Println("PNG画像です")
	}
	return nil
}
