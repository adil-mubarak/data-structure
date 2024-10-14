//auto complete

package main

import "fmt"

type TrieNode struct{
	children map[rune]*TrieNode
	isEnd bool
}

type Trie struct{
	root *TrieNode
}

func NewTrie()*Trie{
	return &Trie{
		root: &TrieNode{children: make(map[rune]*TrieNode)},
	}
}

func(t *Trie)Insert(word string){
	node := t.root
	for _,ch := range word{
		if _,found := node.children[ch]; !found{
			node.children[ch] = &TrieNode{children: map[rune]*TrieNode{}}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie)Search(word string)*TrieNode{
	node := t.root
	for _,ch := range word{
		if _,found := node.children[ch]; !found{
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func(t *Trie)CollectWords(node *TrieNode,prefix string,word *[]string){
	if node == nil{
		return
	}
	if node.isEnd{
		*word = append(*word, prefix)
	}
	for ch,childNode := range node.children{
		t.CollectWords(childNode,prefix+string(ch),word)
	}
}

func(t *Trie)AutoComplete(prefix string)[]string{
	node := t.Search(prefix)
	words := []string{}
	t.CollectWords(node,prefix,&words)
	return words
}

func main() {
	trie := NewTrie()
	trie.Insert("dog")
	trie.Insert("dove")
	trie.Insert("dot")
	trie.Insert("doll")

	words := trie.AutoComplete("do")
	fmt.Println(words)
}