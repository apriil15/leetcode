package main

func main() {

}

// time: O(n)
// space: O(k)
func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int, 0, len(nums)-k+1)
	queue := make([]int, 0, k) // store index in queue

	for i, num := range nums {
		// if first element is out of window -> remove it
		if i >= k && len(queue) > 0 && queue[0] == i-k {
			queue = queue[1:]
		}

		// once num > last element -> remove last element
		for len(queue) > 0 && num > nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}

		queue = append(queue, i)

		// till i reach window -> nums[queue[0]] (will be max)
		if i >= k-1 {
			result = append(result, nums[queue[0]])
		}
	}

	return result
}
