package main

import (
	"cmp"
	"slices"
)

func main() {

}

// time: O(nlogn)
// space: O(n)
func carFleet(target int, position []int, speed []int) int {
	pairs := make([]Pair, 0, len(position))
	for i := 0; i < len(position); i++ {
		pairs = append(pairs, Pair{
			position: position[i],
			speed:    speed[i],
		})
	}

	slices.SortFunc(pairs, func(p1, p2 Pair) int {
		return cmp.Compare(p1.position, p2.position)
	})

	var stack []float32
	for i := len(pairs) - 1; i >= 0; i-- {
		stack = append(stack, float32(target-pairs[i].position)/float32(pairs[i].speed))
		if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2] {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack)
}

type Pair struct {
	position int
	speed    int
}
