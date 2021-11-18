package main

import (
	"github.com/quzhi1/GolangPlayground/clockface/svg"
	"os"
	"time"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
