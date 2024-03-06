# 66. Plus One

## 題目

[Plus One - LeetCode](https://leetcode.com/problems/plus-one/)

## 思路

1. 這題不需要什麼資料結構，很單純的演算法題
2. 主要就是要考慮到`進位`的問題
   - 9999 → 10000
   - 1234 → 1235

## 解法

### Go

```go
package main

import (
	"fmt"
)

func main() {
	result := plusOne([]int{9, 9, 9, 9})
	fmt.Println(result)
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0 // 9 -> 要進位
	}

	// 代表 9999 -> 10000
	newDigits := make([]int, len(digits)+1)
	newDigits[0] = 1
	return newDigits
}
```

### TypeScript

```typescript
// time: O(n)
// space: O(1) / O(n)
function plusOne(digits: number[]): number[] {
  const n = digits.length
  for (let i = n - 1; i >= 0; i--) {
    if (digits[i] < 9) {
      digits[i]++
      return digits
    }
    digits[i] = 0
  }

  // 99 -> 100
  digits.unshift(1)
  return digits
}
```
