package sort

import (
	"github.com/leetcode/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	originArr = []int{3, 1, 4, 5, 2}
	resArr    = [][]int{
		{3, 1}, {3, 2}, {4, 2}, {5, 2},
	}

	originArr1 = []int{6, 8, 4, 3, 2}
	resArr1    = [][]int{
		{6, 4}, {6, 3}, {6, 2}, {8, 4}, {8, 3}, {8, 2}, {4, 3}, {4, 2}, {3, 2},
	}
)

func TestReversePairNormal(t *testing.T) {
	res := ReversePairNormal(originArr)
	assert.Equal(t, resArr, res)

	res = ReversePairNormal(originArr1)
	assert.Equal(t, resArr1, res)
}

func TestReversePair(t *testing.T) {
	for i := 0; i < maxDiff; i++ {
		arr := util.GenRandomArray(arrayLength, maxInt)
		reversePairNormalArr := make([]int, len(arr))
		copy(reversePairNormalArr, arr)
		reversePairArr := ReversePair(arr)
		reversePairNormal := ReversePairNormal(reversePairNormalArr)
		assert.Equal(t, len(reversePairNormal), len(reversePairArr))
	}
}
