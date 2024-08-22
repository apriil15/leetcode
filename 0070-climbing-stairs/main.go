package main

func main() {

}

// time: O(n)
// space: O(1)
func climbStairs(n int) int {
	// n = 1 -> 1
	// n = 2 -> 2
	// n = 3 -> 3
	// n = 4 -> 5
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	first := 1
	second := 2

	var sum int
	for i := 3; i <= n; i++ {
		sum = first + second

		first = second
		second = sum
	}
	return sum
}
