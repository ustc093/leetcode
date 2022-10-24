package sort

func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		value := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if value > arr[j] {
				break
			}
			arr[j+1] = arr[j]
		}

		// insert res arr,index = j
		arr[j+1] = value
	}
}
