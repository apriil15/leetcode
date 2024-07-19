package main

func main() {

}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	digitToLetters := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	var res []string
	var comb string

	var backtrack func(i int)
	backtrack = func(i int) {
		if i == len(digits) {
			res = append(res, comb)
			return
		}

		for _, char := range digitToLetters[digits[i]] {
			comb += char
			backtrack(i + 1)
			comb = comb[:len(comb)-1]
		}
	}

	backtrack(0)
	return res
}
