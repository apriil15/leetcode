# 83. Remove Duplicates from Sorted List

## 題目

[Remove Duplicates from Sorted List - LeetCode](https://leetcode.com/problems/remove-duplicates-from-sorted-list/)

## 思路

解法一

1. 宣告幾個 pointer 去操作 head

2. 遇到已經在 map 裡面的，就把 prev 指向 following

3. 要稍微想一下 pointer 該怎麼移動

解法二

因為題目已經`排序過`，所以不需要用 map

## 解法一 (map)

```go
package main

func main() {

}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	m := make(map[int]interface{})

	prev := head
	current := head
	following := head

	for current != nil {
		following = following.Next

		if _, ok := m[current.Val]; ok {
			prev.Next = following
		} else {
			m[current.Val] = 0
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
```

## 解法二

```go
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
```
