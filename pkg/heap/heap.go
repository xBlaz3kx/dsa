package heap

type Heap struct {
	elements []int
}

func NewHeap() *Heap {
	return &Heap{elements: []int{}}
}

func (h *Heap) Add(element int) {
	for i := 0; i < len(h.elements); i++ {

	}

	// Heapify
}

func (h *Heap) MaxHeap() {

}

func (h *Heap) MinHeap() {

}

func (h *Heap) heapify(array []int, n, index int) {
	largest := index
	leftNode := 2*index + 1
	rightNode := 2*index + 2

	if leftNode < n && array[leftNode] > array[largest] {
		largest = leftNode
	}

	if rightNode < n && array[rightNode] > array[largest] {
		largest = rightNode
	}

	if largest != index {
		tmp := array[index]
		array[index] = array[largest]
		array[largest] = tmp

		h.heapify(array, n, largest)
	}
}
