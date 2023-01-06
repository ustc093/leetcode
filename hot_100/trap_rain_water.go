package hot_100

import "math"

// origin: https://leetcode.cn/problems/trapping-rain-water/
// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6 = 1 + 1 + 2 + 1 + 1

func trap(height []int) int {
	var res int
	if height == nil || len(height) == 0 {
		return res
	}
	leftIndex := 0
	rightIndex := len(height) - 1
	for ; leftIndex < rightIndex; {
		min := math.Min(float64(height[rightIndex]), float64(height[leftIndex]))

		if height[leftIndex] <= height[rightIndex] {
			tmpLeftIndex := leftIndex
			for ;tmpLeftIndex < rightIndex; {
				if height[leftIndex] < height[tmpLeftIndex] {
					break
				} else {
					res += int(min) - height[tmpLeftIndex]
					tmpLeftIndex++
				}
			}
			leftIndex = tmpLeftIndex
		} else {
			tmpRightIndex := rightIndex
			for ;leftIndex < tmpRightIndex; {
				if height[tmpRightIndex] > height[rightIndex] {
					break
				} else {
					res += int(min) - height[tmpRightIndex]
					tmpRightIndex--
				}

			}
			rightIndex = tmpRightIndex
		}
	}

	return res
}
