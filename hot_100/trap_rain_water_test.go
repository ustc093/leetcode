package hot_100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testArray = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	testArrayRes = 6
	testArrayII = []int{4,2,0,3,2,5}
	testArrayResII = 9
)

func TestTrap(t *testing.T) {
	sum := trap(testArray)
	assert.Equal(t, testArrayRes, sum)

	sum = trap(testArrayII)
	assert.Equal(t, testArrayResII, sum)
}
