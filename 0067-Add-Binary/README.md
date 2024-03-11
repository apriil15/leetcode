# 67. Add Binary

## 題目

[Add Binary - LeetCode](https://leetcode.com/problems/add-binary/)

## 思路

1. 這題與 [415. Add Strings](https://leetcode.com/problems/add-strings/) 很類似
2. 差別在十進位與`二進位`

## 解法

```go
func addBinary(a string, b string) string {
	// 取得 maxLength
	maxLength := 0
	if len(a) >= len(b) {
		maxLength = len(a)
	} else {
		maxLength = len(b)
	}

	result := make([]byte, maxLength+1) // +1 預留進位
	carry := byte(0)                    // 進位的數字

	for i := 0; i < maxLength; i++ {
		num1 := byte(0)
		num2 := byte(0)

		// 最後一位依序取出
		if i < len(a) {
			num1 = a[len(a)-1-i] - '0'
		}
		if i < len(b) {
			num2 = b[len(b)-1-i] - '0'
		}

		s := num1 + num2 + carry
		carry = s / 2
		result[maxLength-i] = s%2 + '0'
	}

	if carry > 0 {
		result[0] = '1'
	} else {
		result = result[1:]
	}
	return string(result)
}
```
