package main

import (
	"fmt"
	"time"
)

// UNBUFFERED CHANNEL

//  GOROUTINE SYNCHRONIZATION
// func task(done chan bool) {
// 	defer func() {
// 		done <- true
// 	}()
// 	fmt.Println("Processing...")
// }

func emailSender(emailChan chan string, done chan bool) {
	defer func() {
		done <- true
	}()
	for email := range emailChan {
		fmt.Println("Sending email to : ", email)
		time.Sleep(time.Second)
	}
}

func main() {

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "shiva"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("Received data from channel 1 : ", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("Received data from channel 2 : ", chan2Val)
		}
	}

	// emailChan := make(chan string, 100)

	// done := make(chan bool)

	// go emailSender(emailChan, done)

	// for i := 0; i < 5; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// fmt.Println("All emails are sent")

	// // this is important to close the channel after done their work
	// close(emailChan)
	// <-done

	// emailChan <- "a@gmail.com"
	// emailChan <- "b@gmail.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)

	// done := make(chan bool)

	// go task(done)

	// <-done
}

// SENDING DATA
// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("Processing number : ", num)
// 		time.Sleep(time.Second)
// 	}
// }

// GETTING DATA
// func sum(result chan int, num1, num2 int) {
// 	numResult := num1 + num2
// 	result <- numResult
// }

// func main() {

// 	result := make(chan int)

// 	go sum(result, 4, 5)

// 	res := <-result

// 	fmt.Println(res)

// numChan := make(chan int)

// go processNum(numChan)

// for {
// 	numChan <- rand.Intn(100)
// }
// }

// GOROUTINES FUNCTIONS
// func main() {
// 	messageChannel := make(chan string)

// 	go func() {
// 		messageChannel <- "Hello from goroutine" // this means store the value in the channel
// 	}()

// 	message := <-messageChannel // this means storing the value of the channel in another variale

// 	fmt.Println(message)
// }
