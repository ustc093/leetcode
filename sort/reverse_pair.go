package sort

/**
 * 逆序对
 * 定义：对于一个包含N个非负整数的数组A[1..n]，如果有i < j，且A[i]>A[j]，则称(A[i], A[j])为数组A中的一个逆序对。
 * 例如：数组[3，1，4，5，2]的逆序对有(3,1),(3,2),(4,2),(5,2)，共4个。
 */

// ReversePair T(N) = O(N*logN)
// 利用merge sort
func ReversePair(arr []int) [][]int {

	res := make([][]int, 0)
	reversePairHelper(arr, 0, len(arr)-1, &res)
	return res
}

func reversePairHelper(arr []int, start, end int, res *[][]int) {
	if start >= end {
		return
	}
	mid := start + (end-start)/2
	reversePairHelper(arr, start, mid, res)
	reversePairHelper(arr, mid+1, end, res)

	// handle reverse Pair
	i := start
	j := mid + 1
	k := start
	tmp := make([]int, len(arr))
	for ; i <= mid && j <= end; {
		if arr[i] > arr[j] {
			for x := i; x <= mid; x++ {
				*res = append(*res, []int{arr[x], arr[j]})
			}
			tmp[k] = arr[j]
			j++

		} else {
			tmp[k] = arr[i]
			i++
		}
		k++
	}

	for ; i <= mid; {
		tmp[k] = arr[i]
		k++
		i++
	}

	for ; j <= end; {
		tmp[k] = arr[j]
		k++
		j++
	}

	for ; start <= end; start++ {
		arr[start] = tmp[start]
	}
}

// ReversePairNormal T(N) = O(N^2)
func ReversePairNormal(arr []int) [][]int {
	res := make([][]int, 0)
	n := len(arr)
	for i, valueI := range arr {
		for j := i + 1; j < n; j++ {
			if valueI > arr[j] {
				tmp := []int{valueI, arr[j]}
				res = append(res, tmp)
			}
		}
	}

	return res
}
