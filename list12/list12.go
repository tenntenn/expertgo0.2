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
	f, err := os.Open("sample.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	h := sha256.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return err
	}

	fmt.Printf("%x\n", h.Sum(nil))
	return nil
}
