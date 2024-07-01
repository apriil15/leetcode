package main

func main() {

}

// time: O(n)
// space: O(1)
func threeConsecutiveOdds(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	var count int
	for _, v := range arr {
		if isOdd(v) {
			count++
			if count == 3 {
				return true
			}
		} else {
			count = 0
		}
	}
	return false
}

func isOdd(n int) bool {
	return n%2 == 1
}
