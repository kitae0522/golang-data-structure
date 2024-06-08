package trie

import "fmt"

/*

TODO
1. Insert
2. Search
3. StartsWith
4. AutoComplete

*/

type Node struct {
	child map[rune]*Node
	isEnd bool
}

type Trie struct {
	root *Node
}

func NewNode() *Node {
	return &Node{child: make(map[rune]*Node)}
}

func NewTrie() *Trie {
	return &Trie{root: NewNode()}
}

func (t *Trie) Insert(word string) {
	curNode := t.root
	for _, char := range word {
		if _, ok := curNode.child[char]; !ok {
			curNode.child[char] = NewNode()
		}
		curNode = curNode.child[char]
	}
	curNode.isEnd = true
}

func (t *Trie) Search(word string) bool {
	curNode := t.root
	for _, char := range word {
		if _, ok := curNode.child[char]; !ok {
			return false
		}
		curNode = curNode.child[char]
	}
	return curNode.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	curNode := t.root
	for _, char := range prefix {
		if _, ok := curNode.child[char]; !ok {
			return false
		}
		curNode = curNode.child[char]
	}
	return true
}

func (t *Trie) AutoComplete(prefix string) []string {
	var (
		result  []string
		curNode *Node = t.root
	)
	for _, char := range prefix {
		if _, ok := curNode.child[char]; !ok {
			return nil
		}
		curNode = curNode.child[char]
	}

	t.collectWords(curNode, prefix, &result)
	return result
}

func (t *Trie) collectWords(node *Node, word string, result *[]string) {
	if node.isEnd {
		*result = append(*result, word)
	}
	for char, child := range node.child {
		t.collectWords(child, word+string(char), result)
	}
}

/*
func main() {
	newTrie := NewTrie()
	sampleString := []string{
		"apple", "apex", "banana", "ban", "banish", "brother", "cool",
	}
	for _, word := range sampleString {
		newTrie.Insert(word)
	}

	fmt.Println(newTrie.AutoComplete("ap"))
	fmt.Println(newTrie.AutoComplete("b"))
	fmt.Println(newTrie.AutoComplete("ban"))
}
*/
