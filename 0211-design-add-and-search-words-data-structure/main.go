package main

func main() {

}

type WordDictionary struct {
	root *TrieNode
}

type TrieNode struct {
	children    map[rune]*TrieNode
	isEndOfWord bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: NewTrieNode(),
	}
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children:    make(map[rune]*TrieNode),
		isEndOfWord: false,
	}
}

func (this *WordDictionary) AddWord(word string) {
	node := this.root
	for _, char := range word {
		if _, ok := node.children[char]; !ok {
			node.children[char] = NewTrieNode()
		}
		node = node.children[char]
	}
	node.isEndOfWord = true
}

func (this *WordDictionary) Search(word string) bool {
	return search(this.root, word)
}

//		            口
//		         /  |   \
//		        b   d    m
//		       /    |     \
//		      a     a      a
//		     /      |       \
//	        d       d        d
func search(node *TrieNode, word string) bool {
	for i, char := range word {
		if char == '.' {
			// for word: ".ad", then do recursion for "bad", "dad" and "mad"
			for _, n := range node.children {
				if search(n, word[i+1:]) {
					return true
				}
			}
			return false
		} else {
			if _, ok := node.children[char]; !ok {
				return false
			}
			node = node.children[char]
		}
	}
	return node.isEndOfWord
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
