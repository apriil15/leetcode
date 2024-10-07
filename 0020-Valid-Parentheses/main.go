package main

func main() {

}

// time: O(n)
// space: O(n)
func isValid(s string) bool {
	var stack []byte
	openToClose := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for i := 0; i < len(s); i++ {
		if isOpen(s[i]) {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}

			open := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if openToClose[open] != s[i] {
				return false
			}
		}
	}
	return len(stack) == 0
}

func isOpen(b byte) bool {
	return b == '(' || b == '[' || b == '{'
}
