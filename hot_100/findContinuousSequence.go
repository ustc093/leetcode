package hot_100

import "sort"

/**
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。


示例 1：
输入：target = 9
输出：[[2,3,4],[4,5]]

示例 2：
输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]]

链接：https://leetcode.cn/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof
*/

func findContinuousSequence(target int) [][]int {
	res := make([][]int, 0)
	if target <= 2 {
		return res
	}
	i := target / 2
	j := i - 1
	curr := i
	for j >= 0 && j < i {
		curr += j
		if curr == target {
			tmp := make([]int, 0)
			for z := 0; z < i-j+1; z++ {
				tmp = append(tmp, j+z)
			}
			res = append(res, tmp)
			curr -= i
			i--
		} else if curr > target{
			curr -= i
			i--
		}
		j--
	}
	if target%2 == 1 {
		res = append(res, []int{target / 2, target/2 + 1})
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0]
	})
	return res
}
