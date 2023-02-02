package hot_100

/**
实现pow(x,n)，即计算 x 的 n 次幂函数（即，x^n）。不得使用库函数，同时不需要考虑大数问题。

示例 1：
输入：x = 2.00000, n = 10
输出：1024.00000

示例 2：
输入：x = 2.10000, n = 3
输出：9.26100

示例 3：
输入：x = 2.00000, n = -2
输出：0.25000
解释：2-2 = 1/22 = 1/4 = 0.25

链接：https://leetcode.cn/problems/shu-zhi-de-zheng-shu-ci-fang-lcof
*/
func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / helpMyPow(x, -n)
	}
	return helpMyPow(x, n)
}

func helpMyPow(x float64, n int) float64 {
	res := float64(1)
	for i := n; i > 0; i /= 2 {
		if i%2 == 1 {
			res *= x
		}
		x *= x
	}

	return res
}
