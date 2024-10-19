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
	cars := make([]car, 0, len(position))
	for i := range len(position) {
		cars = append(cars, car{
			position: position[i],
			speed:    speed[i],
		})
	}

	slices.SortFunc(cars, func(c1, c2 car) int {
		return cmp.Compare(c1.position, c2.position)
	})

	var stack []float32
	for i := len(cars) - 1; i >= 0; i-- {
		costTime := float32(target-cars[i].position) / float32(cars[i].speed)
		stack = append(stack, costTime)

		if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2] {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack)
}

type car struct {
	position int
	speed    int
}
