# 118. Pascal's Triangle

## 題目

[Pascal's Triangle - LeetCode](https://leetcode.com/problems/pascals-triangle/)

## 思路

1. 左右兩側都為 1
2. 中間部分為上面兩個的和，並試著推導出規律怎麼用程式表達

```go
// Input: numRows = 5
// Output:
//     [1],
//    [1,1],
//   [1,2,1],
//  [1,3,3,1],
// [1,4,6,4,1]

// result[2][0] = 1
// result[2][1] = result[1][0] + result[1][1] // 試著推導出規律
// result[2][2] = 1
```

## 解法

```go
package main

import (
	"fmt"
)

func main() {
	numRows := 6
	result := generate(numRows)
	fmt.Println(result)
}

func generate(numRows int) [][]int {
	// 宣告好 slice
	result := make([][]int, numRows)
	for i := range result {
		result[i] = make([]int, i+1) // 因為是長度，所以要 +1
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j < i+1; j++ {
			if j == 0 || i == j {
				result[i][j] = 1 // 左右兩側都是 1
			} else {
				result[i][j] = result[i-1][j-1] + result[i-1][j] // 用上面推導出來的規律
			}
		}
	}
	return result
}
```