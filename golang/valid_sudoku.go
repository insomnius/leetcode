package leetcode

import "fmt"

func isValidSudoku(board [][]byte) bool {
	rowsMap := map[string]bool{}
	columnMap := map[string]bool{}
	boxMap := map[string]bool{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == []byte(".")[0] {
				continue
			}

			rowKey := fmt.Sprintf("%s%d", string(board[i][j]), i)
			columnKey := fmt.Sprintf("%s%d", string(board[i][j]), j)
			boxKey := fmt.Sprintf("%s%d%d", string(board[i][j]), i/3, j/3)

			if _, found := rowsMap[rowKey]; found {
				return false
			}

			if _, found := columnMap[columnKey]; found {
				return false
			}

			if _, found := boxMap[boxKey]; found {
				return false
			}

			rowsMap[rowKey] = true
			columnMap[columnKey] = true
			boxMap[boxKey] = true
		}
	}

	return true
}
