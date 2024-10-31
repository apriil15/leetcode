package main

func main() {

}

// time: O(n)
// space: O(1)
func maxProfit(prices []int) int {
	left := 0
	right := 0
	var res int
	for right <= len(prices)-1 {
		if prices[right] <= prices[left] {
			left = right
		}

		res = max(res, prices[right]-prices[left])

		right++
	}
	return res
}
