package main

func main() {

}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//  取得長度
	lengthA := GetLength(headA)
	lengthB := GetLength(headB)

	// 移到同一個起點
	if lengthA > lengthB {
		for i := 0; i < lengthA-lengthB; i++ {
			headA = headA.Next
		}
	} else {
		for i := 0; i < lengthB-lengthA; i++ {
			headB = headB.Next
		}
	}

	// 一一比對
	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 獲得 node 的長度
func GetLength(node *ListNode) int {
	length := 0
	current := node
	for current != nil {
		current = current.Next
		length++
	}
	return length
}
