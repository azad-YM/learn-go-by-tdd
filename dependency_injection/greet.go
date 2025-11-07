package dependency_injection

import (
	"fmt"
	"io"
	"net/http"
)

type Log struct {
	message string
}

func (l *Log) Write(p []byte) (n int, err error) {
	l.message = string(p)
	return len(p), nil
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func GreetController(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World")
}
