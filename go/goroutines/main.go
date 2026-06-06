package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Taks is done : ", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func(i int, w *sync.WaitGroup) {
			task(i, w)
		}(i, &wg)
	}

	wg.Wait()
}
