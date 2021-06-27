# 234. Palindrome Linked List

## 題目

[Palindrome Linked List - LeetCode](https://leetcode.com/problems/palindrome-linked-list/)

## 思路

第一種解法：

1. 遍歷 linked list 把 value 都存進 slice
2. 對 slice 頭尾比對是否對稱

第二種解法：

1. 先找到中間的 node
2. 把後半部 reverse
3. 前後兩部分依序比對

## 解法一

```go
package main

import (
	"fmt"
)

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}
	result := isPalindrome(head)
	fmt.Println(result)
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	// 先把 Val 都存進 slice
	current := head
	s := []int{}
	for current != nil {
		s = append(s, current.Val)
		current = current.Next
	}

	// 前後比對是否對稱
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}
```

## 解法二

follow up: `space complexity O(1)`

```go
package main

import (
	"fmt"
)

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}
	result := isPalindrome(head)
	fmt.Println(result)
}

func isPalindrome(head *ListNode) bool {
	// 用兩個 pointer，來找到中間的 node，目的是 reverse 後半部的 node
	// slow: 一次移動一格
	// fast: 一次移動兩格
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil { // 代表 node 長度為奇數
		slow = slow.Next
	}

	// reverse 後半段 node
	reverseNode := reverse(slow)

	// 前後半段比對
	for head != nil && reverseNode != nil {
		if head.Val != reverseNode.Val {
			return false
		}
		head = head.Next
		reverseNode = reverseNode.Next
	}
	return true
}

func reverse(node *ListNode) *ListNode {
	var prev *ListNode
	current := node
	following := node
	for current != nil {
		following = following.Next
		current.Next = prev
		prev = current
		current = following
	}
	return prev
}

type ListNode struct {
	Val  int
	Next *ListNode
}
```
