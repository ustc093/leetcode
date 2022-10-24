package sort

import (
	"github.com/leetcode/util"
)

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return util.MergeArr(left, right)
}

func MergeSort2(arr []int) {

	helper(arr, 0, len(arr)-1)
}

func helper(arr []int, start, end int) {
	if start >= end {
		return
	}
	mid := start + (end-start)/2
	helper(arr, start, mid)
	helper(arr, mid+1, end)
	util.MergeArr2(arr, start, end)
}
