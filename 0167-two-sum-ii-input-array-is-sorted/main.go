package main

func main() {

}

// time: O(n)
// space: O(1)
func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		if numbers[left]+numbers[right] > target {
			right--
			continue
		}
		if numbers[left]+numbers[right] < target {
			left++
			continue
		}
		return []int{left + 1, right + 1}
	}
	return nil
}
