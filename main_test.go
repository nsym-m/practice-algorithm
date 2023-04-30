package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	trie := newTrie()
	trie.insert("word")
	trie.insert("wheel")
	trie.insert("world")
	trie.insert("hospital")
	trie.insert("mode")

	cases := []struct {
		searchItem string
		expected   string
	}{
		{
			searchItem: "word",
			expected:   "word",
		},
		{
			searchItem: "wo",
			expected:   "wo",
		},
		{
			searchItem: "wh",
			expected:   "wh",
		},
		{
			searchItem: "wor",
			expected:   "wor",
		},
		{
			searchItem: "host",
			expected:   "",
		},
		{
			searchItem: "mode",
			expected:   "mode",
		},
		{
			searchItem: "code",
			expected:   "",
		},
	}

	for _, c := range cases {
		res := trie.match(c.searchItem)
		if res != c.expected {
			t.Errorf("actual:%v expected:%v\n", res, c.expected)
		}
	}
}
