package main

func main() {

}

//         1           2   3   4
//    /    |   \      /|\
//   2     3    4    1 3 4
//  / \   /\    /\
// 3  4  2  4  2 3
// |  |  |  |  | |
// 4  3  4  2  3 2

// time complexity: O(n*n!)
func permute(nums []int) [][]int {
	result := [][]int{}

	var backtrack func(m map[int]struct{}, arr []int)
	backtrack = func(m map[int]struct{}, arr []int) {
		if len(arr) == len(nums) {
			dest := make([]int, 0, len(arr))
			dest = append(dest, arr...)

			result = append(result, dest)
			return
		}

		for i, num := range nums {
			if _, ok := m[i]; ok {
				continue
			}
			m[i] = struct{}{}

			arr = append(arr, num)
			backtrack(m, arr)

			delete(m, i)
			arr = arr[:len(arr)-1]
		}
	}

	backtrack(map[int]struct{}{}, []int{})
	return result
}
