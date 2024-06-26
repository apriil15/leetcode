package main

import (
	"cmp"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

func main() {

}

// use library: https://github.com/emirpasic/gods
type KthLargest struct {
	// if this minHeap has only k size,
	// then minHeap's top -> kth largest
	minHeap *priorityqueue.Queue
	k       int
}

func asc(a, b any) int {
	return cmp.Compare(a.(int), b.(int))
}

func Constructor(k int, nums []int) KthLargest {
	minHeap := priorityqueue.NewWith(asc)
	for _, v := range nums {
		minHeap.Enqueue(v)
		if minHeap.Size() > k {
			minHeap.Dequeue()
		}
	}

	return KthLargest{
		minHeap: minHeap,
		k:       k,
	}
}

// time: O(log n)
func (this *KthLargest) Add(val int) int {
	this.minHeap.Enqueue(val)
	if this.minHeap.Size() > this.k {
		this.minHeap.Dequeue()
	}

	v, _ := this.minHeap.Peek()
	return v.(int)
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
