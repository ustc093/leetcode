package hot_100

import (
	"strings"
)

/**
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数在数组的前半部分，所有偶数在数组的后半部分。


示例：

输入：nums =[1,2,3,4]
输出：[1,3,2,4]
注：[3,1,2,4] 也是正确的答案之一。

链接：https://leetcode.cn/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof
*/

func exchange(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	left := 0
	right := len(nums) - 1
	for left < right {
		for left < len(nums) && nums[left]%2 != 0 {
			left++
		}

		for right >= 0 && nums[right]%2 == 0 {
			right--
		}

		if right > left {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}



	return nums
}

func reverseWords(s string) string {
	res := ""
	sArr := strings.Split(s," ")
	right := len(sArr) - 1
	for right >= 0 {
		if sArr[right] != "" {
			res += " " + sArr[right]
		}
		right--
	}
	return strings.Trim(res, " ")
}
