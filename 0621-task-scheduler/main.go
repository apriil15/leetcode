package main

import (
	"cmp"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

func main() {

}

func leastInterval(tasks []byte, n int) int {
	// map for counter
	counter := make(map[byte]int)
	for _, b := range tasks {
		counter[b]++
	}

	// use map to build a maxHeap
	maxHeap := priorityqueue.NewWith(func(a, b any) int {
		return -cmp.Compare(a.(int), b.(int))
	})
	for _, v := range counter {
		maxHeap.Enqueue(v)
	}

	queue := []Node{}

	var res int
	for maxHeap.Size() > 0 || len(queue) > 0 {
		res++
		if maxHeap.Size() > 0 {
			tmp, _ := maxHeap.Dequeue()
			v := tmp.(int)
			v--

			if v != 0 {
				queue = append(queue, Node{
					count:    v,
					idleTime: res + n,
				})
			}
		}
		if len(queue) > 0 && queue[0].idleTime == res {
			first := queue[0]
			queue = queue[1:]

			maxHeap.Enqueue(first.count)
		}
	}

	return res
}

type Node struct {
	count    int
	idleTime int
}
