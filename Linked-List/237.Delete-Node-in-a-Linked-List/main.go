package main

func main() {

}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
