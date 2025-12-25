package main

import (
	"log"
	"time"

	"github.com/sony/gobreaker"
)

func NewMongoBreaker() *gobreaker.CircuitBreaker {
	settings := gobreaker.Settings{
		Name:        "mongo-db",
		MaxRequests: 2,
		Timeout:     10 * time.Second,

		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures >= 3
		},

		OnStateChange: func(name string, from, to gobreaker.State) {
			log.Printf("Circut %s changed from %v to %v - ", name, from, to)
		},
	}

	return gobreaker.NewCircuitBreaker(settings)
}
