package main

import "slices"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// time: O(n^2)
// space: O(n)
//
// prerequisite:
// preorder: root left right
// inorder: left root right
//
// ideas:
// preorder[0] must be "root", but we don't know the rest of left subtree, right subtree.
// so use preorder[0] to find the root "index" in inorder,
// the left part will be left subtree, right part will be right subtree.
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	mid := slices.Index(inorder, preorder[0])

	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:mid+1], inorder[:mid]),
		Right: buildTree(preorder[mid+1:], inorder[mid+1:]),
	}
}
