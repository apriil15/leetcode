# 24. Swap Nodes in Pairs

## 題目

[Swap Nodes in Pairs - LeetCode](https://leetcode.com/problems/swap-nodes-in-pairs/)

## 思路

1.  可以先從兩個開始想，1 → 2，變成 2 → 1

    這樣就能 reverse 了

    ```go
    following := head.Next
    head.Next = nil
    following.Next = head
    return following
    ```

2.  接下來去想怎麼把兩個 reverse 的組合在一起，可以用 `recursion`

## 解法

```go
package main

func main() {

}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	following := head.Next
	head.Next = swapPairs(head.Next.Next)
	following.Next = head
	return following
}

type ListNode struct {
	Val  int
	Next *ListNode
}
```
