package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	arr     = []int{4, 3, 5, 1, 3, 5, 6, 8}
	sortArr = []int{1, 3, 3, 4, 5, 5, 6, 8}
)

func TestSelectSort(t *testing.T) {
	SelectSort(arr)
	assert.Equal(t, arr, sortArr)
}
