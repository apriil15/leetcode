package main

func main() {

}

// time: O(n)
// space: O(n)
func isValid(s string) bool {
	stack := make([]byte, 0)

	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for i := 0; i < len(s); i++ {
		if isOpenBracket(s[i]) {
			stack = append(stack, s[i])
			continue
		}

		if len(stack) > 0 && stack[len(stack)-1] == m[s[i]] {
			stack = stack[0 : len(stack)-1] // pop last one
		} else {
			return false
		}
	}

	return len(stack) == 0
}

func isOpenBracket(b byte) bool {
	return b == '(' || b == '[' || b == '{'
}
