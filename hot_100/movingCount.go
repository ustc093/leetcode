package hot_100

/**
地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。
一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。
例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？


示例 1：
输入：m = 2, n = 3, k = 1
输出：3

示例 2：
输入：m = 3, n = 1, k = 0
输出：1

链接：https://leetcode.cn/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof
*/

var (
	visited [][]bool
	m1,n1 int
)

func movingCount(m int, n int, k int) int {
	if k < 0 {
		return 0
	}
	m1 = m
	n1 = n
	visited = make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	return helpMovingCount(0, 0, k)
}

func helpMovingCount(i int, j int, k int) int {
	if i >= m1 || j >= n1 || countSum(i)+countSum(j) > k || visited[i][j] {
		return 0
	}
	visited[i][j] = true
	return 1 + helpMovingCount(i+1, j, k) + helpMovingCount(i, j+1, k)
}

func countSum(num int) int {
	var res int
	for num != 0 {
		pop := num % 10
		res += pop
		num /= 10
	}

	return res
}
