package sort

import "github.com/leetcode/util"

func BubbleSort(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j+1] > arr[j] {
				continue
			}
			util.SwapArray(arr, j, j+1)
		}
	}
}
