# 206. Reverse Linked List

## 題目

[Reverse Linked List - LeetCode](https://leetcode.com/problems/reverse-linked-list/)

## 思路

1. 因為要讓 current.Next 指向 prev，原本指向的 node 會找不到，因此需要一個 following 先把原本 current.Next 指向的 node 存起來
2. 指向處理完，prev 移到 current 位置，current 移到 following 位置
3. 最後回傳 prev，原因是上面的 loop 既然結束跳出來，代表 current 是 nil（那當然 following 也也會是 nil）

## 解法

```go
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
```

- 也可以用這個比較精簡的寫法，只宣告一個 node，但我比較喜歡上面的作法，比較有步驟感 XD

```go
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head, head.Next, prev = head.Next, prev, head
	}
	return prev
}
```

## Reference

這篇我覺得解釋地很詳細，推薦

[Reversing a Linked List: Easy as 1, 2, 3](https://medium.com/outco/reversing-a-linked-list-easy-as-1-2-3-560fbffe2088)
