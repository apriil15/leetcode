package main

import "fmt"

func main() {
	a := "11"
	b := "1"
	result := addBinary(a, b)
	fmt.Println(result)
}

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
