# 357. Count Numbers with Unique Digits

## 題目

[Count Numbers with Unique Digits - LeetCode](https://leetcode.com/problems/count-numbers-with-unique-digits/)

## 思路

1. 數學題
2. 規律

   ```
   1 位數 → 10
   2 位數 → 9 * 9
   3 位數 → 9 * 9 * 8
   4 位數 → 9 * 9 * 8 * 7
   5 位數 → 9 * 9 * 8 * 7 * 6
   .
   .
   .
   10 位數 → 9 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1
   11 位數以上都會是 0，因為一定會重複
   ```

3. 以 3 位數舉例，區間為 `100 ~ 999`，不重複的個數為 `9 * 9 * 8`
   - 百位可以填 `9` 個數字（1 ~ 9）
   - 十位可以填 `9` 個數字（原本可以填 10 個數字，但百位用掉 1 個）
   - 個位可以填 `8` 個數字（原本可以填 10 個數字，但百位跟十位都用掉 1 個）

## 解法

```go
package main

import "fmt"

func main() {
	result := countNumbersWithUniqueDigits(2)
	fmt.Println(result)
}

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}

	result := 10
	for i := 1; i < n; i++ { // 從 1 開始是因為 n >= 2
		count := 9
		for j := 0; j < i; j++ {
			count *= 9 - j
		}
		result += count
	}
	return result
}
```
