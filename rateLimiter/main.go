package main

import (
	"fmt"
	"sync"
	"time"
)

// TokenBucket struct
type TokenBucket struct {
	capacity     int
	tokens       chan struct{} // channel to represent tokens
	fillInterval time.Duration
}

// NewTokenBucket creates a new token bucket
func NewTokenBucket(capacity int, fillInterval time.Duration) *TokenBucket {
	tb := &TokenBucket{
		capacity:     capacity,
		tokens:       make(chan struct{}, capacity),
		fillInterval: fillInterval,
	}

	// Fill the bucket initially
	for i := 0; i < capacity; i++ {
		tb.tokens <- struct{}{}
	}

	// Start refill goroutine
	go func() {
		ticker := time.NewTicker(fillInterval)
		for range ticker.C {
			select {
			case tb.tokens <- struct{}{}: // refill a token if not full
			default:
				// bucket full, do nothing
			}
		}
	}()

	return tb
}

// Request simulates sending a request with queuing behavior
func (tb *TokenBucket) Request(id int, wg *sync.WaitGroup) {
	defer wg.Done() // mark this request as done when finished

	fmt.Printf("Request %d: waiting for token...\n", id)
	<-tb.tokens // Wait until token is available
	fmt.Printf("Request %d: processing ✅\n", id)
	time.Sleep(1 * time.Second) // simulate processing
	fmt.Printf("Request %d: done ✅\n", id)
}

func main() {
	tb := NewTokenBucket(3, time.Second)

	var wg sync.WaitGroup

	// Simulate 10 requests arriving quickly
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go tb.Request(i, &wg)
		time.Sleep(300 * time.Millisecond) // new request every 0.3 sec
	}

	// Wait for all requests to finish
	wg.Wait()
	fmt.Println("✅ All requests finished")
}
