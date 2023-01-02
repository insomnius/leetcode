package leetcode

// 289. Game of Life
// According to Wikipedia's article: "The Game of Life, also known simply as Life, is a cellular automaton devised by the British mathematician John Horton Conway in 1970."

// The board is made up of an m x n grid of cells, where each cell has an initial state: live (represented by a 1) or dead (represented by a 0). Each cell interacts with its eight neighbors (horizontal, vertical, diagonal) using the following four rules (taken from the above Wikipedia article):

// Any live cell with fewer than two live neighbors dies as if caused by under-population.
// Any live cell with two or three live neighbors lives on to the next generation.
// Any live cell with more than three live neighbors dies, as if by over-population.
// Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
// The next state is created by applying the above rules simultaneously to every cell in the current state, where births and deaths occur simultaneously. Given the current state of the m x n grid board, return the next state.

// Example 1:
// https://assets.leetcode.com/uploads/2020/12/26/grid1.jpg

// Input: board = [[0,1,0],[0,0,1],[1,1,1],[0,0,0]]
// Output: [[0,0,0],[1,0,1],[0,1,1],[0,1,0]]
// Example 2:
// https://assets.leetcode.com/uploads/2020/12/26/grid2.jpg

// Input: board = [[1,1],[1,0]]
// Output: [[1,1],[1,1]]

// Constraints:

// m == board.length
// n == board[i].length
// 1 <= m, n <= 25
// board[i][j] is 0 or 1.

func gameOfLife(board [][]int) {
	// Golang copy would copy slices as references
	previousStateBoard := [][]int{}
	for k, rows := range board {
		previousStateBoard = append(previousStateBoard, []int{})
		for _, element := range rows {
			previousStateBoard[k] = append(previousStateBoard[k], element)
		}
	}

	for columnK, rows := range board {
		for rowK, element := range rows {
			neighbour := countNeighbour(previousStateBoard, columnK, rowK)
			if element == 0 {
				if neighbour == 3 {
					board[columnK][rowK] = 1
				}
			} else {
				if !(neighbour == 2 || neighbour == 3) {
					board[columnK][rowK] = 0
				}
			}
		}
	}
}

func countNeighbour(previousStateBoard [][]int, columnK, rowK int) int {
	occurrence := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			if columnK+i < 0 || rowK+j < 0 {
				continue
			}

			if columnK+i > len(previousStateBoard)-1 || rowK+j > len(previousStateBoard[columnK])-1 {
				continue
			}

			if v := previousStateBoard[columnK+i][rowK+j]; v == 1 || v == 3 {
				occurrence++
			}
		}
	}

	return occurrence
}

func gameOfLifeAlt(board [][]int) {
	for columnK, rows := range board {
		for rowK, element := range rows {
			neighbour := countNeighbour(board, columnK, rowK)
			if element == 0 {
				if neighbour == 3 {
					board[columnK][rowK] = 2
				}
			} else {
				if !(neighbour == 2 || neighbour == 3) {
					board[columnK][rowK] = 1
				} else {
					board[columnK][rowK] = 3
				}
			}
		}
	}

	for columnK, rows := range board {
		for rowK := range rows {
			if board[columnK][rowK] == 1 {
				board[columnK][rowK] = 0
			} else if board[columnK][rowK] == 2 {
				board[columnK][rowK] = 1
			} else if board[columnK][rowK] == 3 {
				board[columnK][rowK] = 1
			}
		}
	}
}
