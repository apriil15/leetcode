# 237. Delete Node in a Linked List

## 題目

[Delete Node in a Linked List - LeetCode](https://leetcode.com/problems/delete-node-in-a-linked-list/)

## 思路

1. 這題怪異的點是傳入的 node 就是要刪除的 node
2. 解法是把這個 node 的值，變成下個 node 的值

## 解法

```go
package main

func main() {

}

func deleteNode(node *ListNode) {
	*node = *node.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
```

- 另外一個寫法

```go
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
```
