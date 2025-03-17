package rate_limiting

import (
	"context"
	"log"
	"sync"
	"time"
)

type TokenBucket struct {
	// Token refill rate in milliseconds.
	refreshRate int

	// Maximum number of tokens available.
	maxCapacity int

	// Current number of tokens available.
	currentTokens int

	mutex sync.Mutex
}

func NewTokenBucket(refreshRate, capacity int) *TokenBucket {
	return &TokenBucket{
		refreshRate:   refreshRate,
		maxCapacity:   capacity,
		currentTokens: capacity,
		mutex:         sync.Mutex{},
	}
}

// ProcessRequests will process requests sent to the request channel.
func (lb *TokenBucket) ProcessRequests(ctx context.Context, requestChan <-chan Request) {
	for {
		select {
		case <-ctx.Done():
			return
		case req := <-requestChan:
			if !lb.HasAvailableTokens() {
				log.Println("no tokens available, discarding request")
				continue
			}

			log.Printf("processed request %s at %s with content %s \n", req.GetID(), time.Now().Format(time.RFC3339), req.GetContent())
			lb.removeToken()
		}
	}
}

// Start will run a ticker which will add tokens to the bucket.
func (lb *TokenBucket) Start(ctx context.Context) {
	processingRate := time.Duration(lb.refreshRate) * time.Millisecond
	ticker := time.NewTicker(processingRate)

	for {
		select {
		case _ = <-ticker.C:
			lb.addToken()
		case <-ctx.Done():
			return
		}
	}
}

// addToken will add a token if the max capacity is not reached.
func (lb *TokenBucket) addToken() {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()

	if lb.currentTokens < lb.maxCapacity {
		lb.currentTokens++
	}
}

// removeToken will remove a token from the bucket
func (lb *TokenBucket) removeToken() {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()

	if lb.currentTokens > 0 {
		lb.currentTokens--
	}
}

// HasAvailableTokens checks if there are tokens available
func (lb *TokenBucket) HasAvailableTokens() bool {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()

	return lb.currentTokens != 0
}
