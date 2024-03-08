# 372. Super Pow

## 題目

[Super Pow - LeetCode](https://leetcode.com/problems/super-pow/)

## 思路

1. 需要知道這個公式：`ab % k = (a % k) * (b % k) % k`（modular exponentiation）
2. 發現可以這樣展開：

   ```
   a^123 % k = (a^120 % k) * (a^3 % k) % k = (a^12 % k)^10 % k * (a^3 % k) % k
   ```

## 解法

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(superPow(2, []int{1, 1}))
}

const MOD = 1337

// 程式的想法
// 例如：a^123 = ((a^1)^10 * (a^2))^10 * a^3
func superPow(a int, b []int) int {
	a %= MOD
	if a == 0 {
		return 0
	}

	result := powMod(a, b[0])
	for i := 1; i < len(b); i++ {
		// 每次都會把前面的結果"十次方"，再乘上新數字
		result = powMod(result, 10) * powMod(a, b[i]) % MOD
	}
	return result
}

func powMod(a int, b int) int {
	if b == 0 {
		return 1
	}

	result := a
	for i := 1; i < b; i++ {
		result = result * a % MOD
	}
	return result
}
```
