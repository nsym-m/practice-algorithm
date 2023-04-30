package main

import (
	// "encoding/json"
	"fmt"
)

func main() {
	text := "aaa"
	ac := newTrie()
	ac.insert("she")
	ac.insert("his")
	ac.insert("hers")
	ac.insert("hes")
	// fmt.Println("res")
	// fmt.Println("")
	res := ac.match(text)

	fmt.Println(res)
}

type AhoCorasick struct{}

type Trie struct {
	Key      int             `json:"key"`
	Children map[string]Trie `json:"children"`
}

type Res struct {
	Key  int    `json:"key"`
	Text string `json:"text"`
}

var key int

func newTrie() Trie {
	return Trie{
		Key:      0,
		Children: make(map[string]Trie),
	}
}

func NewAhoCorasick(text string) AhoCorasick {
	return AhoCorasick{}
}

func (t Trie) insert(word string) {
	for _, v := range word {
		if node, ok := t.Children[string(v)]; ok {
			t = node
		} else {
			key++
			t.Children[string(v)] = Trie{
				Key:      key,
				Children: make(map[string]Trie),
			}
			t = t.Children[string(v)]
		}
	}
}

func (t Trie) match(word string) string {
	node := t
	// fmt.Println(node)
	// byteArray, err := json.MarshalIndent(node, "", "	")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(string(byteArray))

	var res string
	for _, v := range word {
		if next, ok := node.Children[string(v)]; ok {
			// fmt.Println(node.Key)
			node = next
			res = res + string(v)
		} else {
			// fmt.Println(node.Key)
			return ""
		}
	}
	return res
}
