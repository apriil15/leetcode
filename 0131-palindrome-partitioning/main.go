package main

func main() {

}

func partition(s string) [][]string {
	var res [][]string
	var part []string

	var backtrack func(i int)
	backtrack = func(i int) {
		if i == len(s) {
			dest := make([]string, len(part))
			copy(dest, part)
			res = append(res, dest)
			return
		}

		for j := i; j < len(s); j++ {
			if isPal(s, i, j) {
				part = append(part, s[i:j+1])
				backtrack(j + 1)
				part = part[:len(part)-1]
			}
		}
	}

	backtrack(0)
	return res
}

func isPal(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
