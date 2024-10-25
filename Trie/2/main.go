//suffix trie

package main

import "fmt"

type SuffixTrieNode struct{
	children map[rune]*SuffixTrieNode
	isEnd bool
}

type SuffixTrie struct{
	root *SuffixTrieNode
}

func NewSuffixTrie() *SuffixTrie{
	return &SuffixTrie{
		root: &SuffixTrieNode{children: make(map[rune]*SuffixTrieNode)},
	}
}

func (t *SuffixTrie) InsertSuffic(suffix string){
	node := t.root
	for _,ch := range suffix{
		if _,found := node.children[ch]; !found{
			node.children[ch] = &SuffixTrieNode{children: make(map[rune]*SuffixTrieNode)}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *SuffixTrie) BuildSuffixTrie(s string) {
	for i := range s {
		t.InsertSuffic(s[i:])
	}
}

func(t *SuffixTrie)Search(substring string)bool{
	node := t.root
	for _,ch := range substring{
		if _,found := node.children[ch]; !found{
			return false
		}
		node = node.children[ch]
	}
	return true
}

func main() {
	suffixTrie := NewSuffixTrie()

	suffixTrie.BuildSuffixTrie("banana")

	fmt.Println(suffixTrie.Search("ana"))
	fmt.Println(suffixTrie.Search("nana"))
	fmt.Println(suffixTrie.Search("ban"))
	fmt.Println(suffixTrie.Search("apple"))
}