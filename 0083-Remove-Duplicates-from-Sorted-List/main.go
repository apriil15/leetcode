package main

func main() {

}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// -101 因為題目限制 val 介於 -100~100
	record := -101

	prev := head
	current := head
	following := head

	for current != nil {
		following = following.Next

		if current.Val == record {
			prev.Next = following
		} else {
			record = current.Val
			prev = current
		}
		current = current.Next
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
