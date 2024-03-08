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

// each iterate
// 1 -> 2 -> 3 -> 4 -> 5
// 1    2 -> 3 -> 4 -> 5
// 1 <- 2    3 -> 4 -> 5
// 1 <- 2 <- 3    4 -> 5
// 1 <- 2 <- 3 <- 4    5
// 1 <- 2 <- 3 <- 4 <- 5

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

// time: O(n)
// space: O(n)
func reverseList_recursion(head *ListNode) *ListNode {
	var result *ListNode
	current := head

	return recursion(result, current)
}

func recursion(result, current *ListNode) *ListNode {
	if current == nil {
		return result
	}

	next := current.Next
	current.Next = result

	result = current
	current = next

	return recursion(result, current)
}
