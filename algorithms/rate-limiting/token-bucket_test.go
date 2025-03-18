package rate_limiting

import (
	"context"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {
	ctx, end := context.WithTimeout(context.Background(), 10*time.Second)
	defer end()

	bucket := NewTokenBucket(1000, 3)
	requestChan := make(chan Request, 100)
	defer close(requestChan)

	var wg sync.WaitGroup
	wg.Add(3)

	// Generate requests for the bucket
	go func() {
		defer wg.Done()

		minSleep := 0
		maxSleep := 300
		for {
			select {
			case <-ctx.Done():
				return
			default:
				requestChan <- NewRequest("has token?")
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(maxSleep-minSleep)+minSleep))
			}
		}

	}()

	go func() {
		defer wg.Done()
		bucket.Start(ctx)
	}()

	go func() {
		defer wg.Done()
		bucket.ProcessRequests(ctx, requestChan)
	}()

	wg.Wait()
}
