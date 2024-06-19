package main

func main() {

}

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children    map[rune]*TrieNode
	isEndOfWord bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		root: NewTrieNode(),
	}
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children:    make(map[rune]*TrieNode),
		isEndOfWord: false,
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this.root
	for _, char := range word {
		if _, ok := node.children[char]; !ok {
			node.children[char] = NewTrieNode()
		}
		node = node.children[char]
	}
	node.isEndOfWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.root
	for _, char := range word {
		if _, ok := node.children[char]; !ok {
			return false
		}
		node = node.children[char]
	}
	return node.isEndOfWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, char := range prefix {
		if _, ok := node.children[char]; !ok {
			return false
		}
		node = node.children[char]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
