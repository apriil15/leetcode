package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

// time: O(n)
// space: O(1)
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head

	for {
		if fast != nil && fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
		} else {
			return false
		}
		slow = slow.Next

		if fast == slow {
			return true
		}
	}
}
