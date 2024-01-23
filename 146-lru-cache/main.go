package main

func main() {

}

type (
	LRUCache struct {
		capacity int
		m        map[int]*Node
		left     *Node // dummy node, LRU node would be left.next
		right    *Node // dummy node
	}

	Node struct {
		pre   *Node
		next  *Node
		key   int
		value int
	}
)

func Constructor(capacity int) LRUCache {
	left := &Node{}
	right := &Node{}
	left.next = right
	right.pre = left

	return LRUCache{
		capacity: capacity,
		m:        map[int]*Node{},
		left:     left,
		right:    right,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; !ok {
		return -1
	} else {
		this.remove(node)
		this.insert(node)

		return node.value
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		this.remove(node)
	}

	n := &Node{key: key, value: value}
	this.insert(n)
	this.m[key] = n

	if len(this.m) > this.capacity {
		lru := this.left.next
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

// insert last (before right node)
func (this *LRUCache) insert(node *Node) {
	pre := this.right.pre

	pre.next = node
	node.pre = pre

	node.next = this.right
	this.right.pre = node
}
