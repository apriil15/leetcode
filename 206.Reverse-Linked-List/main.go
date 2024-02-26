package main

import "fmt"

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{2, nil}}
	result := reverseList(head)
	fmt.Println(result)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// time: O(n)
// space: O(1)
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var result *ListNode
	current := head

	for current != nil {
		next := current.Next

		current.Next = result
		result = current

		current = next
	}

	return result
}

// each iterate
// 1 -> 2 -> 3 -> 4 -> 5
// 1    2 -> 3 -> 4 -> 5
// 1 <- 2    3 -> 4 -> 5
// 1 <- 2 <- 3    4 -> 5
// 1 <- 2 <- 3 <- 4    5
// 1 <- 2 <- 3 <- 4 <- 5
