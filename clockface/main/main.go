package main

import (
	"os"
	"time"

	"github.com/quzhi1/GolangPlayground/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
