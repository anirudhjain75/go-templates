package main

import (
	"container/heap"
	"fmt"
)

type Object struct {
	Val []int
	Sum int
}

type ObjectHeap []Object

func (h ObjectHeap) Len() int {
	return len(h)
}

func (h ObjectHeap) Less(i, j int) bool {
	return h[i].Sum < h[j].Sum
}

func (h ObjectHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ObjectHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *ObjectHeap) Push(x interface{}) {
	*h = append(*h, x.(Object))
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {

	t := &ObjectHeap{}

	heap.Init(t)

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			heap.Push(t, Object{
				Val: []int{nums1[i], nums2[j]},
				Sum: nums1[i] + nums2[j],
			})
		}
	}

	fmt.Println(t)

	res := make([][]int, 0, k)

	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(t).(Object).Val)
	}

	return res
}

func main() {
	in1 := []int{1,7,11}
	in2 := []int{2,4,6}
	k := 3
	kSmallestPairs(in1, in2, k)
}


