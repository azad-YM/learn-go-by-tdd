package main

import (
	"os"
	"time"

	"example.com/learn/clockface"
)

func main() {
	time := time.Now()
	clockface.SVGWriter(os.Stdout, time)
}
