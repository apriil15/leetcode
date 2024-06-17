package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// inorder:   left root right
// postorder: left right root
//
// time: O(n)
// space: O(n)
func buildTree(inorder []int, postorder []int) *TreeNode {
	valueToIndex := make(map[int]int)
	for i, v := range inorder {
		valueToIndex[v] = i
	}

	return build(
		inorder, postorder,
		0, len(inorder)-1, 0, len(postorder)-1,
		valueToIndex,
	)
}

func build(
	inorder, postorder []int,
	inStart, inEnd, postStart, postEnd int,
	valueToIndex map[int]int,
) *TreeNode {
	if inStart > inEnd || postStart > postEnd {
		return nil
	}

	mid := valueToIndex[postorder[postEnd]]
	leftCount := mid - inStart

	return &TreeNode{
		Val: postorder[postEnd],
		Left: build(
			inorder, postorder,
			inStart, mid-1, postStart, postStart+leftCount-1,
			valueToIndex,
		),
		Right: build(
			inorder, postorder,
			mid+1, inEnd, postStart+leftCount, postEnd-1,
			valueToIndex,
		),
	}
}
