package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Message struct {
	To, From string
	Body     string
}

func Post(m *Message) (rerr error) {
	pr, pw := io.Pipe()
	go func() {
		defer pw.Close()
		enc := json.NewEncoder(pw)
		err := enc.Encode(m)
		if err != nil {
			rerr = err
		}
	}()
	const url = "http://example.com"
	const contentType = "application/json"
	_, err := http.Post(url, contentType, pr)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run() error {
	err := Post(&Message{To: "hoge", From: "hogera", Body: "hello"})
	if err != nil {
		return err
	}
	fmt.Println("POSTしました")
	return nil
}
