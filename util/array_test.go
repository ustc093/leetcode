package util

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestSwapArray(t *testing.T) {
	arr := []int{1, 1, 4, 5, 6}
	i := 1
	j := 3
	assert.Equal(t, 1, arr[i])
	assert.Equal(t, 5, arr[j])
	SwapArray(arr, i, j)
	assert.Equal(t, 5, arr[i])
	assert.Equal(t, 1, arr[j])
}

func TestMergeArr(t *testing.T) {
	arr := []int{1, 1, 4, 5, 6}
	arr1 := []int{2, 3, 5, 10}
	res := MergeArr(arr, arr1)
	sortArr := append(arr, arr1...)
	sort.Ints(sortArr)
	assert.Equal(t, sortArr, res)
}
