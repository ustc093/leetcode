package hot_100

import (
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	var res string
	if len(nums) == 0 {
		return res
	}
	sort.Slice(nums, func(i, j int) bool {
		strI := strconv.Itoa(nums[i])
		strJ := strconv.Itoa(nums[j])

		parseIntI, _ := strconv.ParseInt(strI+strJ, 10, 64)
		parseIntJ, _ := strconv.ParseInt(strJ+strI, 10, 64)
		return parseIntI < parseIntJ
	})

	for _, num := range nums {
		res += strconv.Itoa(num)
	}
	return res
}
