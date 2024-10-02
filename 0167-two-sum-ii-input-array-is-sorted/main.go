package main

func main() {

}

// time: O(n)
// space: O(1)
func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum > target {
			right--
			continue
		}
		if sum < target {
			left++
			continue
		}

		return []int{left + 1, right + 1}
	}
	return nil
}
