package hot_100

import "math"

func isStraight(nums []int) bool {
	tmp := make([]int, 14)
	min := float64(14)
	max := float64(-1)
	for _, num := range nums {
		if num != 0 {
			if tmp[num] >= 1 {
				return false
			}
			tmp[num]++
			min = math.Min(min, float64(num))
			max = math.Max(max, float64(num))
		}
	}

	if (max - min + 1)  <= 5  {
		return true
	}

	return false
}
