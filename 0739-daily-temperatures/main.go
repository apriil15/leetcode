package main

func main() {

}

// time: O(n)
// space: O(n)
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := []Node{}

	for i := 0; i < len(temperatures); i++ {
		// once we see larger temperature,
		// pop all smaller temperature, and update res
		for len(stack) > 0 && temperatures[i] > stack[len(stack)-1].value {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			res[last.index] = i - last.index
		}

		stack = append(stack, Node{
			index: i,
			value: temperatures[i],
		})
	}

	return res
}

type Node struct {
	index int
	value int
}
