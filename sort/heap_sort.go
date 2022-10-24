package sort

import "github.com/leetcode/util"

func HeapSort(arr []int) {

	if len(arr) <= 1 {
		return
	}

	for index := range arr {
		heapInsert(arr, index)
	}
	n := len(arr)
	for index := range arr {
		heapSize := n - 1 - index
		util.SwapArray(arr, 0, heapSize)
		adjustHeap(arr, 0, heapSize)
	}
}

func heapInsert(arr []int, index int) {
	parent := (index - 1) / 2
	for ; arr[parent] < arr[index]; {
		util.SwapArray(arr, parent, index)
		index = parent
		parent = (index - 1) / 2
	}
}

func adjustHeap(arr []int, parent, n int) {
	leftChild := parent*2 + 1
	for ; leftChild < n; {
		largest := leftChild
		if leftChild+1 < n && arr[leftChild] < arr[leftChild+1] {
			largest = leftChild + 1
		}

		if arr[largest] < arr[parent] {
			largest = parent
		}

		if largest == parent {
			break
		}
		util.SwapArray(arr, largest, parent)
		parent = largest
		leftChild = parent*2 + 1
	}

}
