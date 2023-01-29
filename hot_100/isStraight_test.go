package hot_100

import (
	"testing"
)

func TestIsStraight(t *testing.T) {
	straight := isStraight([]int{3, 0, 6, 2, 4})
	println(straight)
}
