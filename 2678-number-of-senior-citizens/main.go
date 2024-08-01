package main

func main() {

}

func countSeniors(details []string) int {
	var res int
	for _, detail := range details {
		x := detail[11] - '0' // tens digit
		y := detail[12] - '0' // units digit

		if 10*x+y > 60 {
			res++
		}
	}

	return res
}
