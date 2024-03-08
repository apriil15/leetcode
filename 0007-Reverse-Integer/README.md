# 7. Reverse Integer

## 題目

[Reverse Integer - LeetCode](https://leetcode.com/problems/reverse-integer/)

## 思路

- 解法一（轉成字串）
    1. 數字轉字串 → 字串倒過來 → 再轉數字
    2. 要注意`正負數`問題
- 解法二（應該比較好）
    1. 先取得數字是幾位數
    2. 取出每個數字，再乘上 10 的次方
    例如：123 → 依序取出 3、2、1，變成 300、20、1
    3. 要注意最大值、最小值的判斷

## 解法一

```go
package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	fmt.Println(reverse(-123))

}

// 轉成字串的作法
func reverse(x int) int {
	// 如果是負數，就先 * -1 變成正數
	tmp := x
	if x < 0 {
		tmp *= -1
	}

	// reverse
	str := strconv.Itoa(tmp)
	reverseStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reverseStr += string(str[i])
	}
	reverseNumber, _ := strconv.Atoi(reverseStr)

	// 如果一開始有 * -1，要再改回來
	if x < 0 {
		reverseNumber *= -1
	}

	if reverseNumber > math.MaxInt32 || reverseNumber < math.MinInt32 {
		return 0
	}
	return reverseNumber
}
```

## 解法二

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(123))
}

func reverse(x int) int {
	result := 0

	// 得到 x 是幾位數
	tmp := x
	length := 0
	for tmp != 0 {
		tmp /= 10
		length++
	}

	count := 1
	for length != 0 {
		digit := x / count % 10 // 3 -> 2 -> 1

		result += digit * int(math.Pow(10, float64(length)-1))
		count *= 10 // 1 -> 10 -> 100
		length--
	}

	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}
```