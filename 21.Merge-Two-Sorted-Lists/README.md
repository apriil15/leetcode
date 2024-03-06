# 21. Merge Two Sorted Lists

## 題目

[Merge Two Sorted Lists - LeetCode](https://leetcode.com/problems/merge-two-sorted-lists/)

## 思路

1. 思考方式可以用最簡單的方法來想

   假如有 Node1 與 Node2，值分別是 `1` 跟 `2`，兩者要 merge，結果肯定是 `1 → 2`，代表是讓 `Node1 的 next 指向 Node2`，反之。

2. 運用遞迴就會讓數字由小到大一個指向一個

## 解法

```go
package main

func main() {

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if (l1 == nil) && (l2 == nil) {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
```
