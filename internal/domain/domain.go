package domain

import "fmt"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(newVal ...int) {
	for _, v := range newVal {
		h.array = append(h.array, v)
		h.maxHeapifyUp(len(h.array)-1)
		h.Print()
	}
}

func (h *MaxHeap) Extract() int {
	if len(h.array) == 0 {
		fmt.Printf("%s", "Empty Array")
		return -1
	}

	extracted := h.array[0]
	li := len(h.array)-1

	h.array[0] = h.array[li]
	h.array = h.array[:li]

	h.maxHeapifyDown(0)

	return extracted
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	li := len(h.array) -1
	l,r := left(index), right(index)
	childToCompare := 0

	for l <= li {
		if l == li {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l,r = left(index), right(index)
		} else {
			return
		}
	}

}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func (h *MaxHeap) Print() {
	fmt.Println(h)
}

func parent(i int) int {
	return (i-1)/2
}

func left(i int) int {
	return 2*i+1
}

func right(i int) int {
	return 2*i+2
}