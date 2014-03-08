gop
===

Easy way to process jobs in parallel in golang

* goroutine
* return results
* support timeout

Example usage:

```go
package main

import (
  "fmt"
  "strings"
  "time"

  "github.com/haifeng/gop"
)

type MyWorker struct {
  str string
}

func (mw *MyWorker) Perform() interface{} {
  start := time.Now()
  time.Sleep(2 * time.Second)
  r := strings.ToUpper(mw.str)
  fmt.Printf("Last %s\n", time.Now().Sub(start))
  return r
}

func main() {
  workers := []gop.Worker{
    &MyWorker{"first"},
    &MyWorker{"second"},
    &MyWorker{"third"},
  }

  start := time.Now()

  g := gop.New(5 * time.Second)

  results := g.Run(workers)
  for _, result := range results {
    r := result.(string)
    fmt.Println(r)
  }

  fmt.Printf("Done, got %d results\n", len(results))
  fmt.Printf("Total: %s\n", time.Now().Sub(start))
}
```

More [examples](https://github.com/haifeng/gop/examples).