package hot_100

/**
给定一个m x n 二维字符网格board 和一个字符串单词word 。如果word 存在于网格中，返回 true ；否则，返回 false 。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

例如，在下面的 3×4 的矩阵中包含单词 "ABCCED"（单词中的字母已标出）。


示例 1：
输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true

示例 2：
输入：board = [["a","b"],["c","d"]], word = "abcd"
输出：false

链接：https://leetcode.cn/problems/ju-zhen-zhong-de-lu-jing-lcof
*/

func exist(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}
	for i, col := range board {
		for j, _ := range col {
			if helpExist(board, word, i, j) {
				return true
			}
		}
	}

	return false
}

func helpExist(board [][]byte, word string, i, j int) bool {
	if word == "" {
		return true
	}

	m := len(board)
	n := len(board[0])

	if i >= m || i < 0 || j >= n || j < 0 || board[i][j] != word[0]{
		return false
	}

	board[i][j] = '0'
	res := helpExist(board, word[1:], i, j+1) || helpExist(board, word[1:], i, j-1) ||
		helpExist(board, word[1:], i+1, j) || helpExist(board, word[1:], i-1, j)
	board[i][j] = word[0]
	return res
}
