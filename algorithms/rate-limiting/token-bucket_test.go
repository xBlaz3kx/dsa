package rate_limiting

import (
	"context"
	"math/rand"
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {
	ctx, end := context.WithTimeout(context.Background(), 10*time.Second)
	defer end()

	bucket := NewTokenBucket(1000, 3)
	requestChan := make(chan Request, 100)
	defer close(requestChan)

	// Generate requests for the bucket
	go func() {
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

	bucket.Start(ctx)
	bucket.ProcessRequests(ctx, requestChan)
}
