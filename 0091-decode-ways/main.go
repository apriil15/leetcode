package main

func main() {

}

// time: O(n)
// space: O(26) -> O(1)
func numDecodings(s string) int {
	if s[:1] == "0" { // leading 0: "0", "01", ...
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	m := map[string]struct{}{
		"1":  {},
		"2":  {},
		"3":  {},
		"4":  {},
		"5":  {},
		"6":  {},
		"7":  {},
		"8":  {},
		"9":  {},
		"10": {},
		"11": {},
		"12": {},
		"13": {},
		"14": {},
		"15": {},
		"16": {},
		"17": {},
		"18": {},
		"19": {},
		"20": {},
		"21": {},
		"22": {},
		"23": {},
		"24": {},
		"25": {},
		"26": {},
	}

	var decode1 int = 1
	var decode2 int = 1
	for i := 1; i < len(s); i++ {
		includeLastOne := s[i-1 : i+1]
		thisOne := s[i : i+1] // exclude last one
		_, ok1 := m[includeLastOne]
		_, ok2 := m[thisOne]
		if !ok1 && !ok2 {
			return 0
		}

		var tmp int
		if ok1 {
			tmp += decode1
		}
		if ok2 {
			tmp += decode2
		}

		decode1 = decode2
		decode2 = tmp
	}
	return decode2
}
