package main

func main() {

}

type (
	LRUCache struct {
		c    int
		m    map[int]*Node
		head *Node // dummy node, LRU node would be head.next
		tail *Node // dummy node
	}

	Node struct {
		pre   *Node
		next  *Node
		key   int
		value int
	}
)

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.pre = head

	return LRUCache{
		c:    capacity,
		m:    map[int]*Node{},
		head: head,
		tail: tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if n, ok := this.m[key]; ok {
		this.remove(n)
		this.insert(n)

		return n.value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if n, ok := this.m[key]; ok {
		this.remove(n)
	}

	n := &Node{key: key, value: value}
	this.insert(n)
	this.m[key] = n

	if len(this.m) > this.c {
		lru := this.head.next
		this.remove(lru)
		delete(this.m, lru.key)
	}
}

func (this *LRUCache) remove(node *Node) {
	pre := node.pre
	next := node.next

	pre.next = next
	next.pre = pre
}

// insert insert last (before tail node)
func (this *LRUCache) insert(node *Node) {
	pre := this.tail.pre

	pre.next = node
	node.pre = pre

	node.next = this.tail
	this.tail.pre = node
}
