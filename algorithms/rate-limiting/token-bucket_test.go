package rate_limiting

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {
	fmt.Println("------ Token bucket -----")
	ctx, end := context.WithTimeout(context.Background(), 10*time.Second)
	defer end()

	buckit := NewTokenBucket(1000, 3)
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
				requestChan <- NewRequest("is bucket leaking?")
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(maxSleep-minSleep)+minSleep))
			}
		}

	}()

	buckit.Start(ctx)
	buckit.ProcessRequests(ctx, requestChan)
}
