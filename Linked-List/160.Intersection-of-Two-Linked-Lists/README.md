# 160. Intersection of Two Linked Lists

## 題目

[Intersection of Two Linked Lists - LeetCode](https://leetcode.com/problems/intersection-of-two-linked-lists/)

## 思路

解法一

1. 先對一個 node 建 map
2. 比對到交集點就回傳

解法二

Follow up: Could you write a solution that runs in `O(n) time` and use only `O(1) memory`?

1. 先取得兩個 head 分別的長度
2. 取得長度的差，將較長的 head 指向跟較短 head 一樣的起點

   舉例：

   - A 長度為 5，B 長度為 6

   ```
   A:     4 → 1 →
                  8 → 4 → 5
   B: 5 → 6 → 1 →
   ```

   - 因為 B 比較長，大了 1，因此讓 headB = headB.Next

   ```
   A:     4 → 1 →
                  8 → 4 → 5
   B:     6 → 1 →
   ```

3. 一一去比對，如果一樣代表有交集

## 解法一

```go
package main

func main() {

}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	// 先對一個 node 建 map
	m := make(map[*ListNode]int)
	for headA != nil {
		m[headA] = 0
		headA = headA.Next
	}

	// 比對到交集點就回傳
	for headB != nil {
		if _, ok := m[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//  取得長度
	lengthA := GetLength(headA)
	lengthB := GetLength(headB)

	// 移到同一個起點
	if lengthA > lengthB {
		for i := 0; i < lengthA-lengthB; i++ {
			headA = headA.Next
		}
	} else {
		for i := 0; i < lengthB-lengthA; i++ {
			headB = headB.Next
		}
	}

	// 一一比對
	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 獲得 node 的長度
func GetLength(node *ListNode) int {
	length := 0
	current := node
	for current != nil {
		current = current.Next
		length++
	}
	return length
}
```
