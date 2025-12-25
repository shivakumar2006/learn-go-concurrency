package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://github.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Request failed to timeout", err)
		return
	}
	defer resp.Body.Close()

	log.Println("Request successfull", resp.StatusCode)

	// http client with global timeout
	client2 := &http.Client{}

	resp1, err := client2.Get("https://github.com")
	if err != nil {
		log.Println("timeout or errors : ", err)
	}
}
