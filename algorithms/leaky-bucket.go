package algorithms

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/xBlaz3kx/dsa/pkg/queue"
)

type Request struct {
	ID      string
	Content string
}

func NewRequest(content string) Request {
	return Request{
		ID:      uuid.New().String(),
		Content: content,
	}
}

type LeakyBucket struct {
	// Processing rate in requests/second.
	rate int

	// Maximum number of requests that are available.
	capacity int

	// FIFO queue to store and process the requests from
	queue *queue.ThreadSafeQueue[Request]
}

func NewLeakyBucket(maxRate, capacity int) *LeakyBucket {
	return &LeakyBucket{
		rate:     maxRate,
		capacity: capacity,
		queue:    queue.NewThreadSafeQueue[Request](),
	}
}

// AddRequestToBucket Add the request to the queue for processing. Returns an error if a message cannot be processed.
func (lb *LeakyBucket) AddRequestToBucket(req Request) error {
	if lb.rate <= lb.queue.Size() {
		return errors.New("cannot process the request - bucket full")
	}

	lb.queue.Push(req)
	return nil
}

// Start will start processing request with a constant flow rate.
func (lb *LeakyBucket) Start(ctx context.Context) {
	processingRate := time.Duration(lb.rate) / time.Second
	ticker := time.NewTicker(processingRate)

	for {
		select {
		case _ = <-ticker.C:
			lb.processNext()
		case <-ctx.Done():
			return
		}
	}
}

func (lb *LeakyBucket) processNext() {
	req, err := lb.queue.Pop()
	if err != nil {
		log.Println("No requests to process")
		return
	}

	log.Printf("processed request %s at %s with content %s", req.ID, time.Now().Format(time.RFC3339), req.Content)
}
