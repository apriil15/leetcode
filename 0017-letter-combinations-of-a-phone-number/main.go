package main

func main() {

}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	var res []string
	var comb string

	var backtrack func(i int)
	backtrack = func(i int) {
		if i == len(digits) {
			res = append(res, comb)
			return
		}

		for _, str := range digitToLetters(digits[i]) {
			comb += str
			backtrack(i + 1)
			comb = comb[:len(comb)-1]
		}
	}

	backtrack(0)
	return res
}

func digitToLetters(char byte) []string {
	if char == '2' {
		return []string{"a", "b", "c"}
	}
	if char == '3' {
		return []string{"d", "e", "f"}
	}
	if char == '4' {
		return []string{"g", "h", "i"}
	}
	if char == '5' {
		return []string{"j", "k", "l"}
	}
	if char == '6' {
		return []string{"m", "n", "o"}
	}
	if char == '7' {
		return []string{"p", "q", "r", "s"}
	}
	if char == '8' {
		return []string{"t", "u", "v"}
	}
	if char == '9' {
		return []string{"w", "x", "y", "z"}
	}
	return []string{}
}
