package hot_100

import (
	"math"
	"strings"
)

/**

请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。

示例 1:
输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

https://leetcode.cn/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=c0glesv
*/

func lengthOfLongestSubstring(s string) int {
	var res int
	if len(s) <= 1 {
		return len(s)
	}
	var start int

	for index := 1; index < len(s); index++ {
		if strings.Contains(s[start:index], string(s[index])) {
			start += strings.Index(s[start:index], string(s[index])) + 1
		}
		res = int(math.Max(float64(res), float64(index-start+1)))
	}
	return res
}
