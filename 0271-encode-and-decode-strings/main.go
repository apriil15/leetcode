package main

import (
	"strconv"
)

func main() {

}

func encodeAndDecode(input []string) []string {
	e := encode(input)
	return decode(e)
}

// encode with delimiter: len(s)#
func encode(strs []string) string {
	var res string
	for _, str := range strs {
		res += strconv.Itoa(len(str)) + "#" + str
	}
	return res
}

func decode(str string) []string {
	var res []string
	var strLength string
	for i := 0; i < len(str); i++ {
		if string(str[i]) == "#" {
			length, _ := strconv.Atoi(strLength)
			s := str[i+1 : i+1+length]

			res = append(res, s)

			i += length
			strLength = ""
		} else {
			strLength += string(str[i])
		}
	}
	return res
}
