package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/haifeng/gop"
)

type MyWorker struct {
	str      string
	duration time.Duration
}

func (mw *MyWorker) Perform() interface{} {
	start := time.Now()
	time.Sleep(mw.duration)
	r := strings.ToUpper(mw.str)
	fmt.Printf("Last %s\n", time.Now().Sub(start))
	return r
}

func main() {
	workers := []gop.Worker{
		&MyWorker{"first", 200 * time.Millisecond},
		&MyWorker{"second", 300 * time.Millisecond},
		&MyWorker{"third", 1000 * time.Millisecond},
	}

	start := time.Now()

	g := gop.New(500 * time.Millisecond)

	results := g.Run(workers)

	fmt.Printf("Done, got %d results\n", len(results))

	fmt.Printf("Total: %s\n", time.Now().Sub(start))

	time.Sleep(3 * time.Second)
}
