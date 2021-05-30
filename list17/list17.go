package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/tenntenn/exec"
)

func main() {
	vers, err := Vers("github.com/tenntenn/greeting")
	fmt.Println(vers, err)
}

func Vers(module string) ([]string, error) {
	dir, err := os.MkdirTemp("", "vers*")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(dir)
	env := &exec.Env{Dir: dir}
	env.Run("go", "mod", "init", "tmpmodule")
	env.Run("go", "get", module)
	pr, pw := io.Pipe()
	go func() {
		env.Stdout = pw
		env.Run("go", "list", "-m", "-versions", "-json", module)
		pw.Close()
	}()
	var vers struct{ Versions []string }
	err = json.NewDecoder(pr).Decode(&vers)
	if err != nil {
		return nil, err
	}
	err = env.Err()
	if err != nil {
		return nil, err
	}
	return vers.Versions, nil
}
