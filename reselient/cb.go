package main

import (
	"log"
	"net/http"
	"time"

	"github.com/sony/gobreaker"
)

var CB *gobreaker.CircuitBreaker

func InitCircuitBreaker() {
	settings := gobreaker.Settings{
		Name:        "HTTP-Service",
		MaxRequests: 3,
		Interval:    60 * time.Second,
		Timeout:     5 * time.Second,

		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures >= 3
		},

		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			log.Printf("Circuit %s changed from %v to %v", name, from, to)
		},
	}

	CB = gobreaker.NewCircuitBreaker(settings)
}

func main() {
	// ✅ MUST initialize first
	InitCircuitBreaker()

	result, err := CB.Execute(func() (interface{}, error) {
		resp, err := http.Get("https://github.com")
		if err != nil || resp.StatusCode >= 500 {
			return nil, err
		}
		return resp, nil
	})

	if err != nil {
		log.Println("Circuit breaker blocked or request failed:", err)
		return
	}

	log.Println("Circuit breaker running:", result)
}
