package main

func main() {

}

// time: O(n)
// space: O(1)
func maxArea(height []int) int {
	var res int
	left := 0
	right := len(height) - 1
	for left < right {
		v := min(height[left], height[right]) * (right - left)
		res = max(res, v)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}
