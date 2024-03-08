package main

import (
	"fmt"
)

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}
	result := isPalindrome(head)
	fmt.Println(result)
}

func isPalindrome(head *ListNode) bool {
	// 用兩個 pointer，來找到中間的 node，目的是 reverse 後半部的 node
	// slow: 一次移動一格
	// fast: 一次移動兩格
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil { // 代表 node 長度為奇數
		slow = slow.Next
	}

	// reverse 後半段 node
	reverseNode := reverse(slow)

	// 前後半段比對
	for head != nil && reverseNode != nil {
		if head.Val != reverseNode.Val {
			return false
		}
		head = head.Next
		reverseNode = reverseNode.Next
	}
	return true
}

func reverse(node *ListNode) *ListNode {
	var prev *ListNode
	current := node
	following := node
	for current != nil {
		following = following.Next
		current.Next = prev
		prev = current
		current = following
	}
	return prev
}

type ListNode struct {
	Val  int
	Next *ListNode
}
