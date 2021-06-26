package main

import "fmt"

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{2, nil}}
	result := reverseList(head)
	fmt.Println(result)
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	following := head

	for current != nil {
		following = following.Next
		current.Next = prev
		prev = current
		current = following
	}
	return prev
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
