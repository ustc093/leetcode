package util

import (
	"github.com/stretchr/testify/assert"
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
