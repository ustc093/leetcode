package hot_100

import "fmt"

/**
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。

示例 1：

输入：matrix = [
[1,2,3],
[4,5,6],
[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：

输入：matrix =[[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]


链接：https://leetcode.cn/problems/shun-shi-zhen-da-yin-ju-zhen-lcof
*/
func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	var startRow, startCol int
	endRow := len(matrix)
	if endRow == 0 {
		return res
	}

	endCol := len(matrix[0])
	// 0→、1⬇、2←、3⬆
	direct := 0
	for startRow < endRow && startCol < endCol {
		switch direct {
		case 0:
			for i := startCol; i < endCol; i++ {
				res = append(res, matrix[startRow][i])
			}
			startRow++
			direct = 1
		case 1:
			for i := startRow; i < endRow; i++ {
				res = append(res, matrix[i][endCol-1])
			}
			endCol--
			direct = 2
		case 2:
			for i := endCol - 1; i >= startCol; i-- {
				res = append(res, matrix[endRow-1][i])
			}
			endRow--
			direct = 3
		case 3:
			for i := endRow - 1; i >= startRow; i-- {
				res = append(res, matrix[i][startCol])
			}
			startCol++
			direct = 0
		default:
			fmt.Errorf("unknown direct")
		}
	}

	return res
}
