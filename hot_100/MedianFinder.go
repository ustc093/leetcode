package hot_100

type MedianFinder struct {
	Median float64
	Nums   []int
}

func Constructor() MedianFinder {
	return MedianFinder{
		Nums: make([]int, 0),
	}
}

func (m *MedianFinder) AddNum(num int) {
	var index int
	for _, item := range m.Nums {
		if num >= item {
			index++
			continue
		}
	}

	m.Nums = append(m.Nums[:index], append([]int{num}, m.Nums[index:]...)...)
}

func (m *MedianFinder) FindMedian() float64 {
	n := len(m.Nums)
	if n == 0 {
		return float64(0)
	}
	if n%2 == 0 {
		m.Median = float64(m.Nums[(n-1)/2] + m.Nums[n/2]) / 2
	} else {
		m.Median = float64(m.Nums[n/2])
	}
	return m.Median
}
