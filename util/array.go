package util

import "math/rand"

func SwapArray(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func GenRandomArray(length int, maxInt int) []int {
	res := make([]int, length)
	for i := 0; i < length; i++ {
		res[i] = rand.Intn(maxInt)
	}
	return res
}

func MergeArr(arr1 []int, arr2 []int) []int {
	res := make([]int, 0)
	i := 0
	j := 0
	for ; i < len(arr1) && j < len(arr2); {
		if arr1[i] <= arr2[j] {
			res = append(res, arr1[i])
			i++
		} else {
			res = append(res, arr2[j])
			j++
		}
	}

	for ; i < len(arr1); {
		res = append(res, arr1[i])
		i++
	}

	for ; j < len(arr2); {
		res = append(res, arr2[j])
		j++
	}

	return res
}

func MergeArr2(arr []int, start, end int) {
	res := make([]int, len(arr))
	mid := start + (end-start)/2
	i := start
	j := mid + 1
	k := start
	for ; i <= mid && j <= end; {
		if arr[i] <= arr[j] {
			res[k] = arr[i]
			i++
		} else {
			res[k] = arr[j]
			j++
		}
		k++
	}

	for ; i <= mid; {
		res[k] = arr[i]
		i++
		k++
	}

	for ; j <= end; {
		res[k] = arr[j]
		j++
		k++
	}

	for index := start; index <= end; index++ {
		arr[index] = res[index]
	}

}
