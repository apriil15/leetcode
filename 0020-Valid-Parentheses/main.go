package main

func main() {

}

// time: O(n)
// space: O(n)
func isValid(s string) bool {
	stack := make([]byte, 0)

	closeToOpen := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for i := 0; i < len(s); i++ {
		b := s[i]
		if isOpenBracket(b) {
			stack = append(stack, b)
			continue
		}

		if len(stack) > 0 && stack[len(stack)-1] == closeToOpen[b] {
			stack = stack[:len(stack)-1] // pop last one
		} else {
			return false
		}
	}

	return len(stack) == 0
}

func isOpenBracket(b byte) bool {
	return b == '(' || b == '[' || b == '{'
}
