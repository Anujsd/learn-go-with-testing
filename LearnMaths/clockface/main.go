package main

import (
	"os"
	"time"

	clockface "example.com/LearnMaths"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
