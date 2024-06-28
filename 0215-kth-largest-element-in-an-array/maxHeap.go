package main

import (
	"cmp"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

// maxHeap, and Dequeue k times
// time: O((n + k)logn) -> O(nlogn)
// space: O(n)
func findKthLargest_maxHeap(nums []int, k int) int {
	maxHeap := priorityqueue.NewWith(func(a, b any) int {
		return -cmp.Compare(a.(int), b.(int))
	})
	for _, v := range nums {
		maxHeap.Enqueue(v)
	}

	var res int
	for i := 0; i < k; i++ {
		v, _ := maxHeap.Dequeue()
		res = v.(int)
	}
	return res
}
