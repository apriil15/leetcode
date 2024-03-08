# 203. Remove Linked List Elements

## 題目

[Remove Linked List Elements - LeetCode](https://leetcode.com/problems/remove-linked-list-elements/)

## 思路

1. 宣告一個 pointer 指向 head，用這個 pointer 來操作
2. 第一個值在 loop 中先不處理，留在最後做判斷

## 解法

```go
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

	if head.Val == val { // 上面的 loop 其實沒考慮到第一個的值，因此在這邊做處理
		return head.Next
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
```
