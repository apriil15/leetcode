# 234. Palindrome Linked List

## 題目

[Palindrome Linked List - LeetCode](https://leetcode.com/problems/palindrome-linked-list/)

## 思路

1. 遍歷 linked list 把 value 都存進 slice
2. 對 slice 頭尾比對是否對稱

## 解法

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
