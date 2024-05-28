# Data structures and algorithm implementations in Go

In this repository, I will be implementing various data structures and algorithms in Go. This is a work in progress and
I will be adding more implementations as I go along.

All the implementations are tested using unit tests. You can find the tests in the `*_test.go` files. You can run the
tests using the `go test -v ./...` command.

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
  