package hot_100

/**
 * 在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

 * 示例 1:
 * 输入：s = "abaccdeff"
 * 输出：'b'

 * 示例 2:
 * 输入：s = ""
 * 输出：' '

链接：https://leetcode.cn/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof
 */

func firstUniqChar(s string) byte {
	var res [26]int
	for _, char := range s {
		res[char - 'a']++
	}

	for index, char := range s {
		if res[char - 'a'] == 1 {
			return s[index]
		}
	}

	return ' '
}
