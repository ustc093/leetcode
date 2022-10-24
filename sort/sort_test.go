package sort

import (
	"github.com/leetcode/util"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

var (
	maxDiff     = 300000
	maxInt      = 100
	arrayLength = 30
)

func TestSelectSort(t *testing.T) {
	for i := 0; i < maxDiff; i++ {
		arr := util.GenRandomArray(arrayLength, maxInt)
		sortArr := make([]int, len(arr))
		copy(sortArr, arr)
		SelectSort(arr)
		sort.Ints(sortArr)
		assert.Equal(t, sortArr, arr)
	}
}

func TestInsertSort(t *testing.T) {
	for i := 0; i < maxDiff; i++ {
		arr := util.GenRandomArray(arrayLength, maxInt)
		sortArr := make([]int, len(arr))
		copy(sortArr, arr)
		InsertSort(arr)
		sort.Ints(sortArr)
		assert.Equal(t, sortArr, arr)
	}
}

func TestBubbleSort(t *testing.T) {
	for i := 0; i < maxDiff; i++ {
		arr := util.GenRandomArray(arrayLength, maxInt)
		sortArr := make([]int, len(arr))
		copy(sortArr, arr)
		BubbleSort(arr)
		sort.Ints(sortArr)
		assert.Equal(t, sortArr, arr)
	}
}

func TestMergeSort(t *testing.T) {
	for i := 0; i < maxDiff; i++ {
		arr := util.GenRandomArray(arrayLength, maxInt)
		sortArr := make([]int, len(arr))
		copy(sortArr, arr)
		arr = MergeSort(arr)
		sort.Ints(sortArr)
		assert.Equal(t, sortArr, arr)
	}
}

func TestMergeSort2(t *testing.T) {
	for i := 0; i < maxDiff; i++ {
		arr := util.GenRandomArray(arrayLength, maxInt)
		sortArr := make([]int, len(arr))
		copy(sortArr, arr)
		MergeSort2(arr)
		sort.Ints(sortArr)
		assert.Equal(t, sortArr, arr)
	}
}

func TestQuickSort(t *testing.T) {
	for i := 0; i < maxDiff; i++ {
		arr := util.GenRandomArray(5, maxInt)
		sortArr := make([]int, len(arr))
		copy(sortArr, arr)
		QuickSort(arr)
		sort.Ints(sortArr)
		assert.Equal(t, sortArr, arr)
	}
}

func TestHeapSort(t *testing.T) {
	for i := 0; i < maxDiff; i++ {
		arr := util.GenRandomArray(5, maxInt)
		sortArr := make([]int, len(arr))
		copy(sortArr, arr)
		HeapSort(arr)
		sort.Ints(sortArr)
		assert.Equal(t, sortArr, arr)
	}
}
