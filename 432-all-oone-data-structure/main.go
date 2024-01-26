package main

import (
	"fmt"
	"math"
)

func main() {
	a := Constructor()

	a.Inc("a")
	a.Inc("b")
	a.Inc("b")
	a.Inc("b")
	a.Inc("b")
	a.Dec("b")

	fmt.Println(a.GetMaxKey())
	fmt.Println(a.GetMinKey())
}

type AllOne struct {
	keyToCount  map[string]int
	countToNode map[int]*Node
	// head <- ... <- tail
	// dummy node, head.next would be max count, tail.pre would be min count
	head, tail *Node
}

// all keys with the same count
type Node struct {
	pre   *Node
	next  *Node
	count int
	keys  map[string]struct{}
}

func Constructor() AllOne {
	head := &Node{count: math.MaxInt}
	tail := &Node{count: math.MinInt}
	head.next = tail
	tail.pre = head

	return AllOne{
		keyToCount:  make(map[string]int),
		countToNode: make(map[int]*Node),
		head:        head,
		tail:        tail,
	}
}

func (this *AllOne) Inc(key string) {
	// with key
	if count, ok := this.keyToCount[key]; ok {
		node := this.countToNode[count]
		keys := node.keys

		// 4 <- 3 (N keys)
		// 4 <- 3 (1 keys)
		if n, ok := this.countToNode[count+1]; ok {
			n.keys[key] = struct{}{}
			this.keyToCount[key] = n.count

			delete(keys, key)
			if len(keys) == 0 {
				delete(this.countToNode, count)
				this.remove(node)
			}
			return
		}

		// 5 <- _ <- 3 (N keys)
		// 5 <- _ <- 3 (1 key)
		n := &Node{
			count: count + 1,
			keys:  map[string]struct{}{key: {}},
		}
		this.countToNode[n.count] = n
		this.keyToCount[key] = n.count

		this.addAfter(n, node)

		delete(keys, key)
		if len(keys) == 0 {
			delete(this.countToNode, count)
			this.remove(node)
		}
		return
	}

	// without key
	// 1 (N keys) <- tail
	this.keyToCount[key] = 1
	if node, ok := this.countToNode[1]; ok {
		node.keys[key] = struct{}{}
		return
	}

	// _ <- tail
	node := &Node{
		count: 1,
		keys:  map[string]struct{}{key: {}},
	}
	this.countToNode[1] = node

	this.insertLast(node)
}

func (this *AllOne) Dec(key string) {
	count := this.keyToCount[key]
	node := this.countToNode[count]
	keys := node.keys

	// 1 (N keys) -> tail
	// 1 (1 key) -> tail
	if count-1 == 0 {
		delete(this.keyToCount, key)
		delete(keys, key)
		if len(keys) == 0 {
			delete(this.countToNode, count)
			this.remove(node)
		}
		return
	}

	// 4 (N keys) -> 3
	// 4 (1 key) -> 3
	if n, ok := this.countToNode[count-1]; ok {
		delete(keys, key)
		if len(keys) == 0 {
			delete(this.countToNode, count)
			this.remove(node)
		}
		this.keyToCount[key] = count - 1
		n.keys[key] = struct{}{}
		return
	}

	// 4 (N keys) -> __ -> 2
	// 4 (1 key) -> __ -> 2
	n := &Node{
		count: count - 1,
		keys:  map[string]struct{}{key: {}},
	}
	this.countToNode[n.count] = n
	this.keyToCount[key] = n.count

	this.addAfter2(node, n)

	delete(keys, key)
	if len(keys) == 0 {
		delete(this.countToNode, count)
		this.remove(node)
	}
}

func (this *AllOne) GetMaxKey() string {
	if this.head.next == this.tail {
		return ""
	}

	for k := range this.head.next.keys {
		return k
	}

	return ""
}

func (this *AllOne) GetMinKey() string {
	if this.head.next == this.tail {
		return ""
	}

	for k := range this.tail.pre.keys {
		return k
	}

	return ""
}

// insertLast insert node at last (before tail)
func (this *AllOne) insertLast(node *Node) {
	pre := this.tail.pre
	pre.next = node
	node.pre = pre

	node.next = this.tail
	this.tail.pre = node
}

// remove remove node
func (*AllOne) remove(node *Node) {
	pre := node.pre
	next := node.next

	pre.next = next
	next.pre = pre
}

// addAfter add b after a.
// a is new node, b is valid node.
// valid means pre and next are not nil.
func (*AllOne) addAfter(a *Node, b *Node) {
	pre := b.pre

	pre.next = a
	a.pre = pre

	a.next = b
	b.pre = a
}

// addAfter add b after a.
// a is valid node, b is new node.
// valid means pre and next are not nil.
func (*AllOne) addAfter2(a *Node, b *Node) {
	next := a.next

	a.next = b
	b.pre = a

	b.next = next
	next.pre = b
}
