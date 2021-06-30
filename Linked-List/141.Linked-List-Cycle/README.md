# 141. Linked List Cycle

## 題目

[Linked List Cycle - LeetCode](https://leetcode.com/problems/linked-list-cycle/)

## 思路

- 解法一

  1.  宣告一個 map，把每個 node 都存進去

  2.  每個 node 都判斷有沒有在 map 裡面，有則代表 cycle

- 解法二 (better)

  Follow up: O(1) space complexity

  1.  利用 `Floyd Cycle Detection Algorithm`（龜兔賽跑）

  2.  相關題目：[202. Happy Number](../../Hash-Table/202.Happy-Number)

## 解法一

```go
package main

func main() {

}

func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]int)

	current := head
	for current != nil {
		if _, ok := m[current]; ok {
			return true
		}
		m[current] = 0
		current = current.Next
	}
	return false
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

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}
```
