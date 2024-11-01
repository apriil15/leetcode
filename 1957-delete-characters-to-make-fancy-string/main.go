package main

import "strings"

func main() {

}

// time: O(n)
// space: O(n) (if not consider result, O(1))
func makeFancyString(s string) string {
	// use strings.Builder because there would be much concat
	builder := &strings.Builder{}
	builder.WriteByte(s[0])

	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			count = 1
		}

		if count < 3 {
			builder.WriteByte(s[i])
		}
	}
	return builder.String()
}
