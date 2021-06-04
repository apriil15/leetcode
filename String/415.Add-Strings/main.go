package main

import (
	"fmt"
)

func main() {
	num1 := "456"
	num2 := "77"
	result := addStrings(num1, num2)
	fmt.Println(result)
}

func addStrings(num1 string, num2 string) string {
	// 取出最長的 len
	maxLength := 0
	if len(num1) >= len(num2) {
		maxLength = len(num1)
	} else {
		maxLength = len(num2)
	}

	sum := make([]byte, maxLength+1) // 長度 +1 因為相加後可能會進位
	carry := byte(0)                 // 進位的數字

	for i := 0; i < maxLength; i++ {
		a := byte(0)
		b := byte(0)

		if i < len(num1) {
			a = num1[len(num1)-i-1] - '0' // 從個位數開始取出每個數字
		}
		if i < len(num2) {
			b = num2[len(num2)-i-1] - '0'
		}

		s := a + b + carry
		carry = s / 10                 // ÷ 10 得到進位的數字，並在下一輪加上去
		sum[len(sum)-i-1] = s%10 + '0' // % 10 取出個位數，並 +'0' 轉成 ascii
	}

	if carry > 0 {
		sum[0] = '1' // 有進位就補上 1
	} else {
		sum = sum[1:] // 因為一開始 slice 長度有多預留 1，沒進位的話要拿掉
	}

	return string(sum)
}
