package hot_100

import "math"

func minOperations(nums []int, x int) int {
	if x == 0 {
		return 0
	}

	if len(nums) == 0 {
		return -1
	}

	res := -1
	first := -1
	last := -1
	if x >= nums[0] {
		first = minOperations(nums[1:], x - nums[0])
	}

	if x >= nums[len(nums)-1] {
		last = minOperations(nums[:len(nums)-1], x - nums[len(nums)-1])
	}

	if first != -1 && last != -1{
		res = int(math.Min(float64(first), float64(last))) +1
	} else if first != -1{
		res = first + 1
	} else if last != -1{
		res = last +1
	}

	return res
}

func minOperationsII(nums []int, x int) int {
	if x == 0 {
		return 0
	}

	res := math.MaxInt64
	n := len(nums)
	sum := 0
	right := n
	for ; right > 0 && sum < x; {
		right--
		sum += nums[right]
	}

	// 最长后缀和
	if sum < x {
		return -1
	}

	for left := -1; left < n; left++{
		if left != -1 {
			sum += nums[left]
		}

		for ;right < n && sum > x; right++{
			sum -= nums[right]
		}

		if sum > x {
			break
		}

		if sum == x {
			res = int(math.Min(float64(res), float64(left + 1+ n - right)))
		}
	}

	if res == math.MaxInt64 {
		return -1
	}
	return res
}

