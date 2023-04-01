package trie_tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrie_InsertString(t *testing.T) {
	tree := NewTrie()
	tree.InsertWord("abc")
	tree.InsertWord("ab")
	count := tree.SearchWord("abc")
	preCount := tree.PrefixNumber("ab")
	assert.Equal(t, 1, count)
	assert.Equal(t, 2, preCount)
	tree.InsertWord("bc")
	tree.DeleteWord("ab")
	count = tree.SearchWord("abc")
	preCount = tree.PrefixNumber("ab")
	assert.Equal(t, 1, count)
	assert.Equal(t, 1, preCount)
	assert.Equal(t, 0, tree.SearchWord("ab"))

	domainVipMap := map[string]map[string]bool{
		"www.baidu.com": {"8.8.8.8":true,"9.9.9.9":true,"10.10.10.10":true},
		"www.m.baidu.com": {"121.42.23.13":true},
	}
	ss := fmt.Sprintf("domainVipMap=[%+v]", domainVipMap)
	println(ss)

	tt := map[string][]string{
		"11": {"1","2"},
		"113": {"1","2"},
	}
	ttt := fmt.Sprintf("===========tr=[%+v]", tt)
	println(ttt)
}
