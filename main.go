package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	text := "she"
	ac := newAhoCorasick()
	ac.insert("she")
	ac.insert("hers")
	ac.insert("hes")
	res := ac.match(text)

	fmt.Println(res)
}

type Trie struct {
	Key      rune           `json:"key"`
	Children map[rune]*Trie `json:"children"`
	Failure  *Trie          `json:"failure"`
	Output   []string       `json:"output"`
}

type AhoCorasick struct {
	root *Trie
}

func newAhoCorasick() AhoCorasick {
	return AhoCorasick{
		root: newTrie(0),
	}
}

func newTrie(char rune) *Trie {
	return &Trie{
		Key:      char,
		Children: make(map[rune]*Trie),
		Failure:  nil,
		Output:   []string{},
	}
}

func (ac AhoCorasick) insert(word string) {
	node := ac.root
	for _, v := range word {
		if _, ok := node.Children[v]; !ok {
			node.Children[v] = &Trie{
				Key:      v,
				Children: make(map[rune]*Trie),
			}
		}
		node = node.Children[v]
	}
}

func (ac AhoCorasick) match(word string) string {
	node := ac.root
	byteArray, err := json.MarshalIndent(node, "", "	")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))

	var res string
	for _, v := range word {
		if next, ok := node.Children[v]; ok {
			node = next
			res = res + string(v)
		} else {
			// node =
			return ""
		}
	}
	return res
}

func (t AhoCorasick) buildFailure() {

}
