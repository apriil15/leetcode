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

// time: O(n * n!)
func permute(nums []int) [][]int {
	var res [][]int
	var per []int
	seen := make(map[int]struct{}) // seen index

	var backtrack func()
	backtrack = func() {
		if len(per) == len(nums) {
			dest := make([]int, len(per))
			copy(dest, per)
			res = append(res, dest)
			return
		}

		for i := 0; i < len(nums); i++ {
			if _, ok := seen[i]; ok {
				continue
			}

			per = append(per, nums[i])
			seen[i] = struct{}{}

			backtrack()

			per = per[:len(per)-1]
			delete(seen, i)
		}
	}

	backtrack()
	return res
}
