package minHeap

// import "fmt"

type MinHeap []int

func NewMinHeap() *MinHeap {
	h := make(MinHeap, 0)
	return &h
}

func (h MinHeap) GetMin() int {
	return h[0]
}

func (h *MinHeap) RemoveMin() int {
	min := (*h)[0]
	// h[0] = h[len(h)-1]
	(*h)[0] = (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	h.MinHeapify(0)
	return min
}

func (h *MinHeap) MinHeapify(i int) {
	min := i
	left := leftChildIndex(i)
	right := rightChildIndex(i)
	// fmt.Println(min, left, right,  (*h)[left], (*h)[right], (*h)[i], (*h)[left] < (*h)[i])
	if left < len(*h) && (*h)[left] < (*h)[i] {
		// fmt.Println("------")
		min = left
	} else if right < len(*h) && (*h)[right] < (*h)[i] {
		min = right
	}
	// fmt.Println(min)
	if min != i {
		(*h)[min], (*h)[i] = (*h)[i], (*h)[min]
		h.MinHeapify(min)
	}
	return
}

func (h *MinHeap) Insert(w int) {
	*h = append(*h, w)
	i := len(*h) - 1
	for i != 0 {
		if (*h)[i] < (*h)[parentIndex(i)] {
			(*h)[i], (*h)[parentIndex(i)] = (*h)[parentIndex(i)], (*h)[i]
		}
		i = parentIndex(i)
	}
	return
}

func parentIndex(i int) int {
	return (i - 1) / 2
}

func leftChildIndex(i int) int {
	return 2*i + 1
}

func rightChildIndex(i int) int {
	return 2*i + 2
}
