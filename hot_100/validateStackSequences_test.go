package hot_100

import "testing"

func TestValidateStackSequences(t *testing.T) {
	validateStackSequences([]int{1,2,3,4,5}, []int{4,5,3,2,1})
}
