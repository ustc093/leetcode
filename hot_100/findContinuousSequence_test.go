package hot_100

import "testing"

func TestFindContinuousSequence(t *testing.T) {

	sequence := findContinuousSequence(66)
	println(sequence)

	remaining := lastRemaining(10,17)
	println(remaining)
}
