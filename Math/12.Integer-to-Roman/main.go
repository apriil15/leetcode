package main

import (
	"fmt"
	"strings"
)

func main() {
	result := intToRoman(58)
	fmt.Println(result)
}

func intToRoman(num int) string {
	mapping := []Mapping{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var result strings.Builder // 效能較好
	for _, value := range mapping {
		for num >= value.Integer {
			result.WriteString(value.Roman)
			num -= value.Integer
		}
	}
	return result.String()
}

type Mapping struct {
	Integer int
	Roman   string
}
