package main

func main() {

}

// time: O(n)
// space: O(n)
func twoSum(nums []int, target int) []int {
	// key: num
	// value: index
	seen := make(map[int]int)

	for i, num := range nums {
		need := target - num

		if v, ok := seen[need]; ok {
			return []int{v, i}
		}

		seen[num] = i
	}

	return []int{}
}
