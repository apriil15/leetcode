package main

import "math"

func main() {

}

// time: O(n log m) (n: len(piles), m: maxPile)
// space: O(1)
func minEatingSpeed(piles []int, h int) int {
	maxPile := piles[0]
	for _, p := range piles {
		maxPile = max(maxPile, p)
	}

	res := maxPile

	// eating speed between 1 ~ maxPile
	// binary search
	left := 1
	right := maxPile
	for left <= right {
		mid := left + (right-left)/2

		var totalTime int
		for _, p := range piles {
			totalTime += int(math.Ceil(float64(p) / float64(mid)))
		}

		if totalTime <= h {
			right = mid - 1
			res = min(res, mid)
		} else {
			left = mid + 1
		}
	}
	return res
}
