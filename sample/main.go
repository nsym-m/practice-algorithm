package main

import (
	"fmt"
)

type TrieNode struct {
	char     rune
	children map[rune]*TrieNode
	fail     *TrieNode
	output   []string
}

type AhoCorasick struct {
	root *TrieNode
}

func NewTrieNode(char rune) *TrieNode {
	return &TrieNode{
		char:     char,
		children: make(map[rune]*TrieNode),
		fail:     nil,
		output:   []string{},
	}
}

func NewAhoCorasick() *AhoCorasick {
	return &AhoCorasick{
		root: NewTrieNode(0),
	}
}

func (ac *AhoCorasick) AddPattern(pattern string) {
	node := ac.root
	for _, char := range pattern {
		if _, ok := node.children[char]; !ok {
			node.children[char] = NewTrieNode(char)
		}
		node = node.children[char]
	}
	node.output = append(node.output, pattern)
}

func (ac *AhoCorasick) BuildFailLinks() {
	queue := []*TrieNode{}
	for _, child := range ac.root.children {
		queue = append(queue, child)
		child.fail = ac.root
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for char, child := range node.children {
			queue = append(queue, child)
			fail := node.fail
			for fail != nil && fail.children[char] == nil {
				fail = fail.fail
			}
			if fail != nil {
				child.fail = fail.children[char]
			} else {
				child.fail = ac.root
			}
			child.output = append(child.output, child.fail.output...)
		}
	}
}

func (ac *AhoCorasick) Search(text string) [][2]interface{} {
	node := ac.root
	matches := [][2]interface{}{}
	for i, char := range text {
		for node != nil && node.children[char] == nil {
			node = node.fail
		}
		if node != nil {
			node = node.children[char]
			for _, pattern := range node.output {
				matches = append(matches, [2]interface{}{i - len(pattern) + 1, pattern})
			}
		} else {
			node = ac.root
		}
	}
	return matches
}

func main() {
	ahoCorasick := NewAhoCorasick()
	patterns := []string{"he", "she", "his", "hers"}
	for _, pattern := range patterns {
		ahoCorasick.AddPattern(pattern)
	}
	ahoCorasick.BuildFailLinks()

	text := "ahishers"
	matches := ahoCorasick.Search(text)
	fmt.Println(matches) // [[2 his] [1 she] [5 he] [5 hers]]
}
