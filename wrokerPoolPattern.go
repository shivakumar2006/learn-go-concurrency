package main

import (
	"fmt"
	"sync"
)

func main() {
	numJobs := 10
	numWorkers := 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	// for workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			worker(workerID, jobs, results)
		}(i)
	}

	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Printf("Results : %d\n", result)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Workers %d processing job %d\n", id, job)
		results <- job * 2
	}
}

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	numJobs := 10
// 	numWorkers := 3

// 	jobs := make(chan int, numJobs)
// 	results := make(chan int, numJobs)

// 	var wg sync.WaitGroup

// 	// start worker goroutine
// 	for i := 1; i <= numWorkers; i++ {
// 		wg.Add(1)
// 		go func(workerID int) {
// 			defer wg.Done()
// 			worker(workerID, jobs, results)
// 		}(i)
// 	}

// 	// enqueue jobs
// 	for i := 1; i <= numJobs; i++ {
// 		jobs <- i
// 	}
// 	close(jobs)

// 	// wait until all the tasks are completed
// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()

// 	for result := range results {
// 		fmt.Printf("Results : %d\n", result)
// 	}
// }

// func worker(id int, jobs <-chan int, results chan<- int) {
// 	for job := range jobs {
// 		fmt.Printf("Workers : %d processing job no %d\n", id, job)
// 		results <- job * 2
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	numJobs := 10
// 	numWorkers := 3

// 	jobs := make(chan int, numJobs)
// 	results := make(chan int, numJobs)

// 	var wg sync.WaitGroup

// 	// start worker goroutine
// 	for i := 1; i <= numWorkers; i++ {
// 		wg.Add(1)
// 		go func(workerID int) {
// 			defer wg.Done()
// 			worker(workerID, jobs, results)
// 		}(i)
// 	}

// 	// enqueue jobs
// 	for i := 1; i <= numJobs; i++ {
// 		jobs <- i
// 	}
// 	close(jobs)

// 	// wait for all waitgroup to complete the task
// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()

// 	// collect result
// 	for result := range results {
// 		fmt.Printf("results : %d\n", result)
// 	}
// }

// func worker(id int, jobs <-chan int, results chan<- int) {
// 	for job := range jobs {
// 		fmt.Printf("Worker %d procesing job %d\n", id, job)
// 		results <- job * 2
// 	}
// }
