# Data structures and algorithm implementations in Go

In this repository, I will be implementing various data structures and algorithms in Go. This is a work in progress and
I will be adding more implementations as I go along.

All the implementations are tested using unit tests. You can find the tests in the `*_test.go` files. You can run the
tests using the `go test -v ./...` command.

## Data Structures

- [Linked List](structures/linked-list/linkedlist.go) (also thread-safe version)
    - [Singly Linked List](structures/linked-list/linkedlist.go)
- [Stack](structures/stack/stack.go)
- [Queue](structures/queue/queue.go)
- [Binary Search Tree](structures/trees/binary-tree.go) (with traversal and search)
- TODO [Heap](heap.go)
- TODO [Graph](structures/graphs/graphs.go) (with BFS, DFS, topological sort and using adjacency list and matrix
  techniques)

## Algorithms

- [Sorting](algorithms/sorting/sorting.go)
    - [Bubble Sort](algorithms/sorting/bubble-sort.go)
    - [Selection Sort](sorting.go)
    - [Insertion Sort](algorithms/sorting/insertion-sort.go)
    - [Merge Sort](algorithms/sorting/merge-sort.go)
    - [Quick Sort](algorithms/sorting/quick-sort.go)
- [Searching](searching.go)
    - [Linear Search](algorithms/search/linear-search.go)
    - [Binary Search](algorithms/search/binary-search.go)
  