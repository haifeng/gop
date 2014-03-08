package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/haifeng/gop"
)

type MyResult struct {
	status int
	body   string
}

type MyWorker struct {
	url string
}

func (mw *MyWorker) Perform() interface{} {
	start := time.Now()
	r, _ := http.Get(mw.url)
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("Last %s\n", time.Now().Sub(start))
	return MyResult{r.StatusCode, string(body)}
}

func main() {
	workers := []gop.Worker{
		&MyWorker{"http://google.com"},
		&MyWorker{"http://yahoo.com"},
		&MyWorker{"http://baidu.com"},
	}

	start := time.Now()

	g := gop.New(10000 * time.Millisecond)

	results := g.Run(workers)
	for _, result := range results {
		r := result.(MyResult)
		fmt.Println(r.status)
	}

	fmt.Printf("Done, got %d results\n", len(results))
	fmt.Printf("Total %s\n", time.Now().Sub(start))
}
