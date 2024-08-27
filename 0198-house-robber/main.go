package main

func main() {

}

// time: O(n)
// space: O(1)
func rob(nums []int) int {
	rob1 := 0
	rob2 := 0
	for _, num := range nums {
		tmp := max(num+rob1, rob2)
		rob1 = rob2
		rob2 = tmp
	}
	return rob2
}
