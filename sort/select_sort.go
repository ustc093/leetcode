package sort

import "github.com/leetcode/util"

func SelectSort(arr []int) {
	n := len(arr)
	for i, valueI := range arr {
		valueIndex := i
		for j := i + 1; j < n; j++ {
			if valueI > arr[j] {
				valueI = arr[j]
				valueIndex = j
			}
		}
		util.SwapArray(arr, i, valueIndex)
	}
}
