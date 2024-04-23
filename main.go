package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/xBlaz3kx/dsa/algorithms"
	linked_list "github.com/xBlaz3kx/dsa/pkg/linked-list"
	"github.com/xBlaz3kx/dsa/pkg/queue"
	"github.com/xBlaz3kx/dsa/pkg/stack"
	"github.com/xBlaz3kx/dsa/pkg/trees"
)

func main() {
	linkedListMain()

	queueMain()

	stackMain()

	binaryTreeMain()

	sortingMain()

	leakyBucket()

	tokenBucket()
}

func leakyBucket() {
	fmt.Println("------ Leaky bucket -----")
	ctx, end := context.WithTimeout(context.Background(), 10*time.Second)
	defer end()

	buckit := algorithms.NewLeakyBucket(10, 3)

	// Generate requests for the bucket
	go func() {
		minSleep := 0
		maxSleep := 500
		for {
			select {
			case <-ctx.Done():
				return
			default:
				// Add random delay to publishing
				err := buckit.AddRequestToBucket(algorithms.NewRequest("is bucket leaking?"))
				if err == nil {
					log.Println("Request scheduled for processing")
				}

				time.Sleep(time.Millisecond * time.Duration(rand.Intn(maxSleep-minSleep)+minSleep))
			}
		}

	}()

	// Start will block the execution until the context is done
	buckit.Start(ctx)
}

func tokenBucket() {
	fmt.Println("------ Token bucket -----")
	ctx, end := context.WithTimeout(context.Background(), 10*time.Second)
	defer end()

	buckit := algorithms.NewTokenBucket(1000, 10)
	requestChan := make(chan algorithms.Request, 100)

	// Generate requests for the bucket
	go func() {
		minSleep := 0
		maxSleep := 500
		for {
			select {
			case <-ctx.Done():
				return
			default:
				requestChan <- algorithms.NewRequest("is bucket leaking?")
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(maxSleep-minSleep)+minSleep))
			}
		}

	}()

	buckit.Start(ctx)
	buckit.ProcessRequests(ctx, requestChan)
}

func linkedListMain() {
	fmt.Println("------ SinglyLinkedList -----")
	list := linked_list.NewLinkedList[int]()
	list.AddElement(2)
	list.AddElement(3)
	list.AddElement(4)
	list.AddElement(5)
	list.AddElement(19)
	list.AddElement(12)
	list.AddElement(11)

	list.GetElements()

	element, err := list.FindElement(11)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(element.Value)

	_, err = list.FindElement(10)
	if err != nil {
		fmt.Println(err)
	}
}

func queueMain() {
	fmt.Println("------ Queue -----")
	q := queue.NewQueue[string]()
	q.Push("123")
	q.Push("124443")
	q.Push("1234")
	q.Push("4444")
	q.Push("44313")

	fmt.Println(q)

	for !q.IsEmpty() {
		val, err := q.Pop()
		if err != nil {
			continue
		}
		fmt.Println(val)
	}
}

func stackMain() {
	fmt.Println("------ Stack -----")
	st := stack.NewStack[int]()

	st.Push(123)
	st.Push(1234)
	st.Push(12444)
	st.Push(1244233)

	fmt.Println(st)

	for !st.IsEmpty() {
		pop, err := st.Pop()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(pop)
	}
}

func binaryTreeMain() {
	fmt.Println("------ Binary tree -----")
	tree := trees.NewBinaryTree()
	tree.Insert(1111)
	tree.Insert(1)
	tree.Insert(212)
	tree.Insert(2)
	tree.Insert(33)
	tree.Insert(3)
	tree.Insert(51)
	tree.Insert(123)

	tree.Print()

	fmt.Println(tree.FindDFS(3111))

	fmt.Println(tree.FindBFS(123))
}
