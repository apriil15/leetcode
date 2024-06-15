package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// time: O(n)
// space: O(n)
//
// preorder: root left right
// inorder: left root right
func buildTree_optimize(preorder []int, inorder []int) *TreeNode {
	inorderValueToIndex := make(map[int]int)
	for i, v := range inorder {
		inorderValueToIndex[v] = i
	}

	return build(
		preorder, inorder,
		0, len(preorder)-1, 0, len(inorder)-1,
		inorderValueToIndex,
	)
}

func build(
	preorder, inorder []int,
	preStart, preEnd, inStart, inEnd int, // index
	inorderValueToIndex map[int]int,
) *TreeNode {
	if preStart > preEnd || inStart > inEnd {
		return nil
	}

	mid := inorderValueToIndex[preorder[preStart]] // root index
	leftCount := mid - inStart

	return &TreeNode{
		Val: preorder[preStart],
		Left: build(
			preorder, inorder,
			preStart+1, preStart+leftCount, inStart, mid-1,
			inorderValueToIndex,
		),
		Right: build(
			preorder, inorder,
			preStart+leftCount+1, preEnd, mid+1, inEnd,
			inorderValueToIndex,
		),
	}
}
