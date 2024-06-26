package rate_limiting

import (
	"context"
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestLeakyBucket(t *testing.T) {
	ctx, end := context.WithTimeout(context.Background(), 10*time.Second)
	defer end()

	buckit := NewLeakyBucket(1000, 5)

	// Generate requests for the bucket
	go func() {
		minSleep := 0
		maxSleep := 300

		for {
			select {
			case <-ctx.Done():
				return
			default:
				// Add random delay to publishing
				err := buckit.AddRequestToBucket(NewRequest("is bucket leaking?"))
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
	buckit.Start(ctx)
}
