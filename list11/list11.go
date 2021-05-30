package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run() error {
	f, err := os.Create("sample.txt")
	if err != nil {
		return err
	}

	h := sha256.New()
	w := io.MultiWriter(f, h)

	// ファイル書き出しと同時にハッシュ値も求める
	_, err = io.WriteString(w, "hello")
	if err != nil {
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}
