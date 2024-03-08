package main

func main() {

}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	current := head

	for current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next // 指向下下個，等同把下一個刪掉的概念
		} else {
			current = current.Next
		}
	}

	if head.Val == val { // 上面的 loop 其實沒考慮到第一個，因此在這邊做處理
		return head.Next
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
