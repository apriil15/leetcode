package main

import (
	"cmp"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

func main() {

}

// minHeap with k count
// time: O((n log k)
// space: O(k)
func findKthLargest(nums []int, k int) int {
	minHeap := priorityqueue.NewWith(func(a, b any) int {
		return cmp.Compare(a.(int), b.(int))
	})
	for _, v := range nums {
		minHeap.Enqueue(v)
		if minHeap.Size() > k {
			minHeap.Dequeue()
		}
	}

	res, _ := minHeap.Peek()
	return res.(int)
}
