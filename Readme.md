# Data structures and algorithms (implemented in Go)

In this repository, I will be implementing various data structures and algorithms using Go.
This is a work in progress and I will be adding more implementations as I learn them.

All the implementations are tested using unit tests. You can find the tests in the `*_test.go` files.

## Data Structures

- [Linked List](structures/linked-list) (with thread-safe version)
    - [Singly Linked List](structures/linked-list/linkedlist.go)
- [Stack](structures/stack/stack.go) (with thread-safe version)
- [Queue](structures/queue) (with thread-safe version)
- [Binary Search Tree](structures/trees/binary-tree.go) (with traversal and search)
- TODO [Heap](heap.go)
- [Graph](structures/graphs) (with BFS, DFS, topological sort, using both adjacency list and matrix
  techniques)

## Algorithms

- [Sorting](algorithms/sorting)
    - [Bubble Sort](algorithms/sorting/bubble-sort.go)
    - TODO [Selection Sort](sorting.go)
    - [Insertion Sort](algorithms/sorting/insertion-sort.go)
    - [Merge Sort](algorithms/sorting/merge-sort.go)
    - [Quick Sort](algorithms/sorting/quick-sort.go)
- [Searching](algorithms/search)
    - [Linear Search](algorithms/search/linear-search.go)
    - [Binary Search](algorithms/search/binary-search.go)
- TODO [Sliding Window](algorithms/sliding-window.go)
- [Rate Limiting](algorithms/rate-limiting)
  - [Leaky Bucket](algorithms/rate-limiting/leaky-bucket.go)
  - [Token Bucket](algorithms/rate-limiting/token-bucket.go)

## Running unit tests

You can run the tests using the `go test -v ./...` command.

```bash
$ go test -v ./...
```

## License

This project is licensed under the MIT License - see the [LICENSE](License.md) file for details.
