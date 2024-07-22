package main

import "slices"

func main() {

}

// time: O(n logn)
// space: O(n)
func sortPeople(names []string, heights []int) []string {
	heightToNames := make(map[int][]string)
	for i, v := range heights {
		heightToNames[v] = append(heightToNames[v], names[i])
	}

	slices.Sort(heights) // ASC

	var res []string
	for i := len(heights) - 1; i >= 0; i-- { // from last one
		res = append(res, heightToNames[heights[i]]...)
	}
	return res
}
