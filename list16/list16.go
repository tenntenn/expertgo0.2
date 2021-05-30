package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var r REPL
	r.Exec(os.Stdout, "1 + 1")
}

type REPL struct{ stmts []string }

func (r *REPL) Exec(w io.Writer, expr string) error {
	file, err := os.CreateTemp("", "repl*.go")
	if err != nil {
		return err
	}

	const src = `package main;import"fmt"; func main(){%s;fmt.Println(%s)}`
	fmt.Fprintf(file, src, strings.Join(r.stmts, ";"), expr)
	file.Close()
	defer os.RemoveAll(file.Name())
	// Goのコードとして実行
	cmd := exec.Command("go", "run", file.Name())
	cmd.Stdout, cmd.Stderr = w, io.Discard
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
