# 2. Add Two Numbers

## 題目

[Add Two Numbers - LeetCode](https://leetcode.com/problems/add-two-numbers/)

## 思路

1. 其實就是很單純把兩個 Linked List 裡面的值加起來，要記得進位。

## 解法

```go
package main

func main() {

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 取得 node 長度
	lengthA := GetLength(l1)
	lengthB := GetLength(l2)

	// 兩個 pointer 分別指向兩個 node
	var longer *ListNode
	var shorter *ListNode
	if lengthA >= lengthB {
		longer = l1
		shorter = l2
	} else {
		longer = l2
		shorter = l1
	}

	// 因為 longer 最後一定會指向 nil，如果最後一位有進位的話，就需要一個 prev 的 pointer
	var prev *ListNode

	// 進位
	carry := 0

	for shorter != nil {
		digit := (shorter.Val + longer.Val + carry) % 10
		carry = (shorter.Val + longer.Val + carry) / 10
		longer.Val = digit

		prev = longer
		shorter = shorter.Next
		longer = longer.Next
	}
	for longer != nil {
		digit := (longer.Val + carry) % 10
		carry = (longer.Val + carry) / 10
		longer.Val = digit

		prev = longer
		longer = longer.Next
	}

	// 有進位的話，補上最後一位數
	if carry > 0 {
		prev.Next = &ListNode{1, nil}
	}

	if lengthA >= lengthB {
		return l1
	}
	return l2
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 取得 node 長度
func GetLength(node *ListNode) int {
	current := node
	length := 0
	for current != nil {
		current = current.Next
		length++
	}
	return length
}
```
