package rate_limiting

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/xBlaz3kx/dsa/structures/queue"
)

type leakyBucketEntry struct {
	ID          string
	Content     string
	SubmittedAt time.Time
}

type LeakyBucket struct {
	// Processing rate in requests/second.
	rate int

	// Maximum number of requests that are available.
	capacity int

	// FIFO queue to store and process the requests from
	queue *queue.ThreadSafeQueue[leakyBucketEntry]
}

func NewLeakyBucket(maxRate, capacity int) *LeakyBucket {
	return &LeakyBucket{
		rate:     maxRate,
		capacity: capacity,
		queue:    queue.NewThreadSafeQueue[leakyBucketEntry](),
	}
}

// AddRequestToBucket Add the request to the queue for processing. Returns an error if a message cannot be processed.
func (lb *LeakyBucket) AddRequestToBucket(req Request) error {
	if lb.capacity <= lb.queue.Size() {
		return errors.New("cannot process the request - bucket full")
	}

	entry := leakyBucketEntry{
		ID:          req.GetID(),
		Content:     req.GetContent(),
		SubmittedAt: time.Now(),
	}
	lb.queue.Push(entry)
	return nil
}

// Start will start processing request with a constant flow rate.
func (lb *LeakyBucket) Start(ctx context.Context) {
	processingRate := time.Duration(lb.rate) * time.Millisecond
	ticker := time.NewTicker(processingRate)

	for {
		select {
		case _ = <-ticker.C:
			lb.processNext()
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}

// processNext will process the next request from the queue.
func (lb *LeakyBucket) processNext() {
	req, err := lb.queue.Pop()
	if err != nil {
		log.Println("No requests to process")
		return
	}

	submittedAt := req.SubmittedAt.Format(time.RFC3339)
	processedAt := time.Now().Format(time.RFC3339)

	log.Printf("processed request %s (submitted at: %s) at %s with content \"%s\"", req.ID, submittedAt, processedAt, req.Content)
}
