# 121. Best Time to Buy and Sell Stock

## 題目

[Best Time to Buy and Sell Stock - LeetCode](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/)

## 思路

1. 有看到更低的價格 → 變更 min
2. 價差變大時 → 變更 max

## 解法

```go
package main

import (
	"fmt"
)

func main() {
	result := maxProfit([]int{7, 6, 5, 4, 3, 2, 1})
	fmt.Println(result)
}

func maxProfit(prices []int) int {
	min := prices[0]
	max := 0

	for _, price := range prices {
		if price < min {
			min = price
		} else if (price - min) > max {
			max = price - min
		}
	}
	return max
}
```