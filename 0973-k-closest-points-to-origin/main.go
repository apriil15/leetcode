package main

import (
	"cmp"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

func main() {

}

// time: O(n log n)
// space: O(n)
func kClosest(points [][]int, k int) [][]int {
	if len(points) == 1 || len(points) == k {
		return points
	}

	// minHeap heapify with comparator: a^2 + b^2
	minHeap := priorityqueue.NewWith(func(a, b any) int {
		p1 := a.([]int)
		p2 := b.([]int)

		return cmp.Compare(
			p1[0]*p1[0]+p1[1]*p1[1],
			p2[0]*p2[0]+p2[1]*p2[1],
		)
	})
	for _, v := range points {
		minHeap.Enqueue(v)
	}

	// Dequeue() k count, and append into res
	var res [][]int
	for i := 0; i < k; i++ {
		v, _ := minHeap.Dequeue()
		res = append(res, v.([]int))
	}
	return res
}
