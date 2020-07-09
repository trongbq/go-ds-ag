package sorting

import (
	"errors"
	"math"
)

// Start from 0
// Parent at position k
// Left child at position 2(k+1)-1
// Right child at position 2(k+1)
// Parent > left and right child
type heap struct {
	n int // Position of the left most empty spot
	q []int
}

// Insert at the left most the the open spot, which is n+1, then bubble up to keep heap's characteristic
func (h *heap) insert(a int) error {
	if h.n == cap(h.q) {
		return errors.New("Error: heap overflow")
	}

	h.q[h.n] = a
	bubbleUp(h, h.n)
	h.n++

	return nil
}

// Withraw min value from heap
// Move right most child to replace the min position
// To keep heap characteristic, bubble down
func (h *heap) extractMin() (int, error) {
	if h.n == 0 {
		return -1, errors.New("Nothing left on heap")
	}

	min := h.q[0]
	h.q[0] = h.q[h.n-1]
	h.n--
	bubbleDown(h, 0)

	return min, nil
}

func bubbleUp(h *heap, k int) {
	kp := parent(k)
	if kp == -1 {
		// Root of heap, nothing to do here
		return
	}
	if h.q[k] < h.q[kp] {
		h.q[k], h.q[kp] = h.q[kp], h.q[k]
		bubbleUp(h, kp)
	}
}

func bubbleDown(h *heap, k int) {
	if k >= h.n {
		return
	}
	minIndex := k
	lcIndex := leftChild(k)
	// Compare with left child
	if lcIndex < h.n && h.q[lcIndex] < h.q[minIndex] {
		minIndex = lcIndex
	}
	// Compare with right child
	if lcIndex+1 < h.n && h.q[lcIndex+1] < h.q[minIndex] {
		minIndex = lcIndex + 1
	}
	if minIndex != k {
		// There is a child that is less than value at k, let swap and bubble down
		h.q[k], h.q[minIndex] = h.q[minIndex], h.q[k]
		bubbleDown(h, minIndex)
	}
}

// Parent position of child at position k is ceil(k/2)-1
func parent(k int) int {
	return int(math.Ceil(float64(k)/2.0) - 1)
}

func leftChild(k int) int {
	return 2*(k+1) - 1
}

func initHeap(arr []int) (*heap, error) {
	s := len(arr)
	h := heap{
		n: 0,
		q: make([]int, s, s),
	}
	for i := 0; i < s; i++ {
		err := h.insert(arr[i])
		if err != nil {
			return nil, err
		}
	}
	return &h, nil
}

func HeapSort(arr []int) {
	h, err := initHeap(arr)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(arr); i++ {
		v, err := h.extractMin()
		if err != nil {
			panic(err)
		}
		arr[i] = v
	}
}
