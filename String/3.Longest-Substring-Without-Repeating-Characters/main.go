package main

import (
	"fmt"
	"math"
)

func main() {
	s := "dvdf"
	result := lengthOfLongestSubstring(s)
	fmt.Println(result)
}

func lengthOfLongestSubstring(s string) int {
	max := 0
	arr := []rune{}

	for _, r := range s {
		for i := 0; i < len(arr); i++ {
			if r == arr[i] {
				arr = arr[i+1:]
				break
			}
		}
		arr = append(arr, r)
		max = int(math.Max(float64(max), float64(len(arr))))
	}
	return max
}
