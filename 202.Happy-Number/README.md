# 202. Happy Number

## 題目

[Happy Number - LeetCode](https://leetcode.com/problems/happy-number/)

## 思路

1. 這題用到一個比較特殊的演算法：`Floyd Cycle Detection Algorithm`（龜兔賽跑）

- 如果是一個循環，那麼兔子、烏龜一定會在某一點遇到雙方

![](Floyd-Cycle-Detection-Algorithm-1.png)

- 如果兔子在終點前沒跟烏龜會合，則不是一個循環

![](Floyd-Cycle-Detection-Algorithm-2.png)

2. 另外要知道的常識為：

- 取最後一個位數用 `%` → `123 % 10 = 3`

- 位數想要往左移動用 `/` → `123 / 10 = 12`

## 解法

```go
package main

import (
	"fmt"
)

func main() {
	n := 19
	result := isHappy(n)
	fmt.Println(result)
}

func isHappy(n int) bool {
	x := n
	y := n
	for {
		x = getDigitSum(x)              // 做一次
		y = getDigitSum(getDigitSum(y)) // 做兩次
		if x == 1 || y == 1 {
			return true
		}
		if x == y {
			return false
		}
	}
}

// 取得每個位數平方的和
// 123 -> 1 + 4 + 9 = 14
func getDigitSum(n int) int {
	var digits []int
	for n != 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	var sum int
	for _, v := range digits {
		sum += v * v
	}
	return sum
}
```

## Reference

[探索 Floyd Cycle Detection Algorithm](https://medium.com/@orionssl/%E6%8E%A2%E7%B4%A2-floyd-cycle-detection-algorithm-934cdd05beb9)
