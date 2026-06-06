package main

import (
	"fmt"
	"sync"
)

type Post struct {
	Views int
	mu    sync.Mutex
}

func (p *Post) inc(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		p.mu.Unlock()
	}()

	p.mu.Lock()
	p.Views += 1
}

func main() {
	var wg sync.WaitGroup

	myPost := Post{Views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}

	wg.Wait()

	fmt.Println(myPost.Views)
}
