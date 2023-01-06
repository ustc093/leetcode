package hot_100

/**
在一个 n * m 的二维数组中，每一行都按照从左到右非递减的顺序排序，每一列都按照从上到下非递减的顺序排序。
请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。


示例:

现有矩阵 matrix 如下：

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]

给定 target=5，返回true。

给定target=20，返回false。

链接：https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof
*/

func findNumberIn2DArray(matrix [][]int, target int) bool {
	var res bool
	if len(matrix) == 0 {
		return res
	}

	colEnd := len(matrix[0]) - 1
	rowEnd := len(matrix) - 1
	rowStart := 0
	colStart := 0
	for ; colStart <= colEnd && rowStart <= rowEnd; {

		if target == matrix[rowStart][colStart] || target == matrix[rowEnd][colEnd]{
			return true
		}

		if target > matrix[rowStart][colEnd] {
			rowStart++
		} else if target < matrix[rowStart][colEnd] {
			colEnd--
		}

		if target > matrix[rowEnd][colStart] {
			colStart++
		} else if target < matrix[rowEnd][colStart] {
			rowEnd--
		} else {
			return true
		}

	}

	return res
}
