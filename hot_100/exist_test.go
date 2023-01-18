package hot_100

import "testing"

func TestExist(t *testing.T) {
	board := make([][]byte, 1)
	board[0] = make([]byte, 1)
	board[0][0] = 'a'
	exist(board, "a")

}

