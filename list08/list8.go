package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

func NewPNG(r io.Reader) (io.Reader, error) {
	magicnum := []byte{137, 80, 78, 71}
	buf := make([]byte, len(magicnum))
	if _, err := io.ReadFull(r, buf); err != nil {
		return nil, err
	}
	if !bytes.Equal(magicnum, buf) {
		return nil, errors.New("PNG画像ではありません")
	}
	pngImg := io.MultiReader(bytes.NewReader(magicnum), r)
	return pngImg, nil
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

	_, err = NewPNG(f)
	if err != nil {
		return err
	}
	fmt.Println("PNG画像です")
	return nil
}
