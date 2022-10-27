package sort

func RadixSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	maxDigit := 0
	maxInt := 0
	for _, value := range arr {
		if value > maxInt {
			maxInt = value
		}
	}

	for ; maxInt > 0; {
		maxDigit++
		maxInt /= 10
	}

	num := 1
	for i := 0; i < maxDigit; i++ {
		var bucket [10][]int
		tmp := make([]int, 0)
		for _, value := range arr {
			n := value / num % 10
			bucket[n] = append(bucket[n], value)
		}

		for j := 0; j < 10; j++ {
			tmp = append(tmp, bucket[j]...)
		}

		for x,v := range tmp {
			arr[x] = v
		}
		num *= 10
	}
}
