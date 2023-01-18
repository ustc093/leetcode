package hot_100

import "testing"

func TestExchange(t *testing.T) {
	res := exchange([]int{1,3,5})
	println(res)

	res = exchange([]int{2,4,6})
	println(res)
}

func TestReverseWords(t *testing.T) {
	words := reverseWords("a good   example")
	println(words)
}
