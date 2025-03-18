package rate_limiting

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestLeakyBucket(t *testing.T) {
	ctx, end := context.WithTimeout(context.Background(), 10*time.Second)
	defer end()

	bucket := NewLeakyBucket(1000, 5)

	var wg sync.WaitGroup
	wg.Add(2)

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
				// Add random delay to publishing
				err := bucket.AddRequestToBucket(NewRequest("is bucket leaking?"))
				if err == nil {
					log.Println("Request scheduled for processing")
				} else {
					log.Println("Request denied")
				}

				time.Sleep(time.Millisecond * time.Duration(rand.Intn(maxSleep-minSleep)+minSleep))
			}
		}
	}()

	// Start will block the execution until the context is done
	go func() {
		defer wg.Done()
		bucket.Start(ctx)
	}()

	wg.Wait()
}
