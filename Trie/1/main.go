//basic trie

package main

import "fmt"

type TrieNode struct{
	children map[rune]*TrieNode
	isEnd bool
}

type Trie struct{
	root *TrieNode
}

func NewTrie() *Trie{
	return &Trie{
		root: &TrieNode{children: make(map[rune]*TrieNode)},
	}
}

func (t *Trie)Insert(word string){
	node := t.root
	for _,ch := range word{
		if _,found := node.children[ch]; !found{
			node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie)Search(word string)bool{
	node := t.root
	for _,ch := range word{
		if _,found := node.children[ch]; !found{
			return false
		}
		node = node.children[ch]
	}
	return node.isEnd
}

func(t *Trie)StartsWith(prefix string)bool{
	node := t.root
	for _,ch := range prefix{
		if _,found := node.children[ch]; !found{
			return false
		}
		node = node.children[ch]
	}
	return true
}

func main() {
	trie := NewTrie()

	trie.Insert("hello")

	fmt.Println(trie.Search("hello"))
	fmt.Println(trie.Search("adil"))

	fmt.Println(trie.StartsWith("hel"))
	fmt.Println(trie.StartsWith("hi"))
}