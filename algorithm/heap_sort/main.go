package main

import "fmt"

type Heap struct {
	data []int
	size int
}

/*
	h(Heap) 堆
	parent	要调整的父亲节点
*/
func adjustHead(h Heap, parent int) {
	left := 2*parent + 1
	right := 2*parent + 2

	maxVal := parent
	if left < h.size && h.data[maxVal] > h.data[left] {
		maxVal = left
	}
	if right < h.size && h.data[maxVal] > h.data[right] {
		maxVal = right
	}
	if maxVal != parent {
		h.data[parent], h.data[maxVal] = h.data[maxVal], h.data[parent]
		adjustHead(h, maxVal)
	}
}

func createHeap(nums []int) Heap {
	h := Heap{}
	h.data = nums
	h.size = len(nums)
	for i := h.size / 2; i >= 0; i-- {
		adjustHead(h, i)
	}
	return h
}

func HeapSort(nums []int) {
	h := createHeap(nums)
	for h.size > 0 {
		h.data[0], h.data[h.size-1] = h.data[h.size-1], h.data[0]
		h.size--
		adjustHead(h, 0)
	}
}

func main() {
	arr := []int{3, 5, 12, 6, 98, 9, 43, 0, 11}
	HeapSort(arr)
	fmt.Println(arr)
}
