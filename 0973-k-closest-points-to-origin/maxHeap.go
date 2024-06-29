package main

import (
	"cmp"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

// maxHeap with k count
// time: O((n+k)logk) -> O(nlogk)
// space: O(k)
func kClosest_maxHeap(points [][]int, k int) [][]int {
	pq := priorityqueue.NewWith(func(a, b any) int {
		p1 := a.([]int)
		p2 := b.([]int)

		return -cmp.Compare(
			p1[0]*p1[0]+p1[1]*p1[1],
			p2[0]*p2[0]+p2[1]*p2[1],
		)
	})
	for _, p := range points {
		pq.Enqueue(p)
		if pq.Size() > k {
			pq.Dequeue()
		}
	}

	var res [][]int
	for i := 0; i < k; i++ {
		p, _ := pq.Dequeue()
		res = append(res, p.([]int))
	}
	return res
}
