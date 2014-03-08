package gop

import "time"

type Worker interface {
	Perform() interface{}
}

type Gop struct {
	timeout time.Duration
}

func New(timeout time.Duration) *Gop {
	return &Gop{timeout}
}

func (self *Gop) Run(workers []Worker) (results []interface{}) {
	c := make(chan interface{})

	for _, w := range workers {
		go func(wkr Worker) { c <- wkr.Perform() }(w)
	}

	n := len(workers)
	for i := 0; i < n; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-time.After(self.timeout):
			return
		}
	}
	return
}
