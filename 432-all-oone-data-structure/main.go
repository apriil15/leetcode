package main

import (
	"math"
)

func main() {

}

type AllOne struct {
	keyToCount  map[string]int
	countToNode map[int]*Node
	head        *Node // dummy node, head.next would be max count
	tail        *Node // dummy node, tail.pre would be min count
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
		keyToCount:  map[string]int{},
		countToNode: map[int]*Node{},
		head:        head,
		tail:        tail,
	}
}

func (this *AllOne) Inc(key string) {
	if count, ok := this.keyToCount[key]; ok {
		if node, ok := this.countToNode[count+1]; ok {
			node.keys[key] = struct{}{}
			this.keyToCount[key] = node.count

			originNode := this.countToNode[count]
			delete(originNode.keys, key)
			if len(originNode.keys) == 0 {
				// delete this node
				delete(this.countToNode, count)

				pre := originNode.pre
				next := originNode.next

				pre.next = next
				next.pre = pre
			}
			return
		}

		n := &Node{
			count: count + 1,
			keys:  map[string]struct{}{key: {}},
		}
		this.countToNode[n.count] = n
		this.keyToCount[key] = count + 1

		originNode := this.countToNode[count]
		pre := originNode.pre

		pre.next = n
		n.pre = pre

		n.next = originNode
		originNode.pre = n

		delete(originNode.keys, key)
		if len(originNode.keys) == 0 {
			// delete this node
			delete(this.countToNode, count)

			pre := originNode.pre
			next := originNode.next

			pre.next = next
			next.pre = pre
		}
		return
	}

	this.keyToCount[key] = 1
	if node, ok := this.countToNode[1]; ok {
		node.keys[key] = struct{}{}
		return
	}

	node := &Node{
		count: 1,
		keys:  map[string]struct{}{key: {}},
	}
	this.countToNode[1] = node

	this.insertLast(node)
}

func (this *AllOne) Dec(key string) {

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

func (this *AllOne) insertLast(node *Node) {
	pre := this.tail.pre
	pre.next = node
	node.pre = pre

	node.next = this.tail
	this.tail.pre = node
}
