package main

import "fmt"

func main() {
	head := &ListNode{1, &ListNode{2, nil}}
	result := swapPairs(head)
	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	following := head.Next
	head.Next = swapPairs(head.Next.Next)
	following.Next = head
	return following
}

type ListNode struct {
	Val  int
	Next *ListNode
}
