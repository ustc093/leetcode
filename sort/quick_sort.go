package sort

import "github.com/leetcode/util"

func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	index := helperQuickSort(arr)
	QuickSort(arr[:index])
	QuickSort(arr[index+1:])
}

func helperQuickSort(arr []int) int {

	flag := arr[0]
	i := 0
	j := len(arr) - 1
	for ; i <= j; {
		if flag >= arr[i] {
			i++
			continue
		}
		util.SwapArray(arr, i, j)
		j--
	}
	util.SwapArray(arr, 0, j)

	return j
}
