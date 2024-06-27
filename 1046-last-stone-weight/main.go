package main

import (
	"cmp"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

func main() {

}

// time: O(nlogn)
// space: O(n)
func lastStoneWeight(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}

	maxHeap := priorityqueue.NewWith(func(a, b any) int {
		return -cmp.Compare(a.(int), b.(int))
	})
	for _, v := range stones {
		maxHeap.Enqueue(v)
	}

	// follow game rule
	for maxHeap.Size() > 1 {
		tmp1, _ := maxHeap.Dequeue()
		tmp2, _ := maxHeap.Dequeue()
		a := tmp1.(int)
		b := tmp2.(int)

		if a != b {
			// a must be larger than b
			maxHeap.Enqueue(a - b)
		}
	}

	if maxHeap.Size() == 0 {
		return 0
	}

	// only one
	res, _ := maxHeap.Peek()
	return res.(int)
}
