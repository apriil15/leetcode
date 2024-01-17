package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := []string{"hello", "world", "i", "am", "apriil"}
	e := encode(input)
	fmt.Println(e) // 5#hello 5#world 1#i 2#am 6#apriil

	output := decode(e)
	fmt.Println(output)
}

func encode(strs []string) string {
	result := ""
	for _, str := range strs {
		result += strconv.Itoa(len(str)) + "#" + str
	}

	return result
}

// 5#hello 5#world 1#i 2#am 6#apriil
func decode(str string) []string {
	result := []string{}

	for i := 0; i < len(str); i++ {
		// find str length
		strLength := ""
		for string(str[i]) != "#" {
			strLength += string(str[i])
			i++
		}
		length, _ := strconv.Atoi(strLength)

		// use str length to count str
		tmpStr := ""
		for j := 0; j < length; j++ {
			i++
			tmpStr += string(str[i])
		}

		result = append(result, tmpStr)
	}

	return result
}
