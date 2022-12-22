package trie_tree

import (
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
}
