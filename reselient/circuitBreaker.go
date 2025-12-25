package main

import (
	"errors"
	"log"
	"net/http"
	"sync"
	"time"
)

type State int

const (
	Closed State = iota
	Open
	HalfOpen
)

var ErrCircuitOpen = errors.New("Circuit open")

type CircuitBreaker struct {
	mu sync.Mutex

	state State

	failureCount int
	successCount int

	failureThreshold int
	successThreshold int

	openTimeout time.Duration
	lastFailure time.Time
}

func NewCircuitBreaker(failureThreshold int, successThreshold int, openTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:            Closed,
		failureThreshold: failureThreshold,
		successThreshold: successThreshold,
		openTimeout:      openTimeout,
	}
}

func (cb *CircuitBreaker) Execute(fn func() error) error {
	cb.mu.Lock()

	// open -> reject request
	if cb.state == Open {
		if time.Since(cb.lastFailure) > cb.openTimeout {
			cb.state = HalfOpen
			cb.successCount = 0
		} else {
			cb.mu.Unlock()
			return ErrCircuitOpen
		}
	}

	cb.mu.Unlock()

	// execute actual func
	err := fn()

	cb.mu.Lock()
	defer cb.mu.Unlock()

	if err != nil {
		cb.onFailure()
		return err
	}

	cb.onSuccess()
	return nil
}

func (cb *CircuitBreaker) onFailure() {
	cb.failureCount++
	cb.successCount = 0

	if cb.failureCount >= cb.failureThreshold {
		cb.state = Open
		cb.lastFailure = time.Now()
	}
}

func (cb *CircuitBreaker) onSuccess() {
	if cb.state == HalfOpen {
		cb.successCount++

		if cb.successCount >= cb.successThreshold {
			cb.reset()
		}
		return
	}

	cb.failureCount = 0
}

func (cb *CircuitBreaker) reset() {
	cb.state = Closed
	cb.failureCount = 0
	cb.successCount = 0
}

func main() {
	cb := NewCircuitBreaker(
		3,
		2,
		5*time.Second,
	)

	err := cb.Execute(func() error {
		resp, err := http.Get("https://slow-service.com")
		if err != nil || resp.StatusCode >= 500 {
			return errors.New("service failed")
		}
		return nil
	})

	if err == ErrCircuitOpen {
		log.Println("Service temporarily unavailable")
	}

	log.Println("service is closed")
}
