package main

import "fmt"

func main() {
	a := "aba"
	b := "adc"
	result := findLUSlength(a, b)
	fmt.Println(result)
}

func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	if len(a) > len(b) {
		return len(a)
	}
	return len(b)
}
