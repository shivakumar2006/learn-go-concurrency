// // concurrency in go using go routines

package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	score := []int{0}

	wg.Add(3)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("TWO R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()

	fmt.Println("All the goroutine fucntion runs successfully;")
	fmt.Println(score)
}

// waitgroup and mutex

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"sync"
// )

// var wg sync.WaitGroup
// var mut sync.Mutex

// var response = []string{"text"}

// func main() {

// 	websites := []string{
// 		"https://go.dev",
// 		"https://google.com",
// 		"https://fb.com",
// 		"https://github.com",
// 	}

// 	for _, web := range websites {
// 		wg.Add(1)
// 		go getStatusCode(web)
// 	}

// 	wg.Wait()

// 	fmt.Println("websites successfully run")
// 	fmt.Println(response)
// }

// func getStatusCode(endpoint string) {
// 	defer wg.Done()

// 	res, err := http.Get(endpoint)

// 	if err != nil {
// 		fmt.Println("OOPS, something went wrong...")
// 	} else {
// 		mut.Lock()
// 		response = append(response, endpoint)
// 		mut.Unlock()
// 		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"sync"
// )

// var wg sync.WaitGroup //Pointer
// var mut sync.Mutex    // Pointer

// var signal = []string{"text\n"}

// func main() {

// 	websites := []string{
// 		"https://go.dev",
// 		"https://google.com",
// 		"https://fb.com",
// 		"https://github.com",
// 	}

// 	for _, web := range websites {
// 		wg.Add(1)
// 		go getStatusCode(web)
// 	}
// 	wg.Wait()

// 	fmt.Println("All websites run successfully")
// 	fmt.Println(signal)
// }

// func getStatusCode(endpoint string) {
// 	defer wg.Done()

// 	res, err := http.Get(endpoint)

// 	if err != nil {
// 		fmt.Println("OOPS, something went wrong with the websites")
// 	} else {
// 		mut.Lock()

// 		signal = append(signal, endpoint)

// 		mut.Unlock()
// 		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"sync"
// )

// var wg sync.WaitGroup //Pointer

// func main() {
// 	websiteList := []string{
// 		"https://go.dev",
// 		"https://google.com",
// 		"https://fb.com",
// 		"https://github.com",
// 	}

// 	for _, web := range websiteList {
// 		go getStatusCode(web)
// 		wg.Add(1)
// 	}
// 	wg.Wait()
// }

// func getStatusCode(endpoint string) {
// 	defer wg.Done()

// 	res, err := http.Get(endpoint)

// 	if err != nil {
// 		fmt.Println("OOPS!, something went wrong with the endpoints...")
// 	} else {
// 		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
// 	}

// }

// Goroutines

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	go greetings("hello")
// 	go greetings("kumar")
// 	go greetings("shiva")

// 	time.Sleep(100 * time.Millisecond)
// }

// func greetings(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	go greetings("hello")
// 	go greetings("shiva")
// 	go greetings("Kumar")

// 	time.Sleep(100 * time.Millisecond)
// }

// func greetings(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }
