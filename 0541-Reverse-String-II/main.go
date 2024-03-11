package main

import "fmt"

func main() {
	s := "abcdefg"
	k := 2
	result := reverseStr(s, k)
	fmt.Println(result)
}

func reverseStr(s string, k int) string {
	arr := []byte(s)
	length := len(arr)

	// 因為每 2k 個，只有其中的前面 k 個要 reverse，因此用 bool 當作開關
	shouldReverse := true

	for i := 0; i < length; i += k {
		if shouldReverse {
			end := i + k - 1
			if end > length-1 { // 如果超出最後一個，就改成最後一個
				end = length - 1
			}
			reverse(arr, i, end)
		}
		shouldReverse = !shouldReverse
	}
	return string(arr)
}

func reverse(arr []byte, start int, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start] // swap
		start++
		end--
	}
}
