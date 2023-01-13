package hot_100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	res := lengthOfLongestSubstring("bbtablud")
	assert.Equal(t, 6, res)

	res = lengthOfLongestSubstring("bbbbb")
	assert.Equal(t, 1, res)

	res = lengthOfLongestSubstring("abcabcbb")
	assert.Equal(t, 3, res)

	res = lengthOfLongestSubstring("pwwkew")
	assert.Equal(t, 3, res)

	res = lengthOfLongestSubstring("a")
	assert.Equal(t, 1, res)

	res = lengthOfLongestSubstring("")
	assert.Equal(t, 0, res)



}
