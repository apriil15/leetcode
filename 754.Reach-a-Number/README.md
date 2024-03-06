# 754. Reach a Number

## 題目

[Reach a Number - LeetCode](https://leetcode.com/problems/reach-a-number/)

## 思路

1. 要先觀察出規律

   ```
   0
   -1 1
   -3 1 -1 3
   -6 0 -2 4 -4 2 0 6
   -10 -2 -4 4 -6 2 0 8 -8 0 6 -2 -4 4 2 10
   ----------------------------------------
   排序，把重複的去掉 -> 發現會對稱
   0
   1
   -3 -1 1 3
   -6 -4 -2 0 2 4 6
   -10 -8 -6 -4 -2 0 2 4 6 8 10
   ----------------------------------------
   因此只需看單邊，且最大值如果是奇數，則都會是奇數；偶數就都會是偶數
   0
   1
   1 3
   0 2 4 6
   0 2 4 6 8 10
   1 3 5 7 9 11 13 15
   ```

2. 要先找到 `≥ target` 的那個最大值，再用最大值跟 target 的差，來找到 step

   例如 target 為 9，≥ 9 的最大值為 10，10 跟 9 的差為 1（奇數），要繼續找到下一個最大值 15（差為偶數）

## 解法

```go
package main

func main() {

}

func reachNumber(target int) int {
	// 因為是對稱的，所以負數轉成正數，只看單邊
	if target < 0 {
		target *= -1
	}

	step := 0
	sum := 0

	for sum < target {
		step++
		sum += step
	}
	for (sum-target)%2 == 1 { // 如果差是奇數的話，就要繼續找
		step++
		sum += step
	}
	return step
}
```
