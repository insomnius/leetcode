package leetcode

import (
	"fmt"
	"testing"
)

func TestGameOfLife(t *testing.T) {
	t.Log("Test: game of life")
	t.Run("Example 1", func(t *testing.T) {
		board := [][]int{
			{0, 1, 0},
			{0, 0, 1},
			{1, 1, 1},
			{0, 0, 0},
		}
		expected := [][]int{
			{0, 0, 0},
			{1, 0, 1},
			{0, 1, 1},
			{0, 1, 0},
		}
		gameOfLife(board)
		for i := range board {
			for j := range board[i] {
				if board[i][j] != expected[i][j] {
					t.Errorf("Expected %d, got %d", expected[i][j], board[i][j])
				}
			}
		}
	})
}

func TestGameOfLifeAlt(t *testing.T) {
	t.Log("Test: game of life")
	t.Run("Example 1", func(t *testing.T) {
		board := [][]int{
			{0, 1, 0},
			{0, 0, 1},
			{1, 1, 1},
			{0, 0, 0},
		}
		expected := [][]int{
			{0, 0, 0},
			{1, 0, 1},
			{0, 1, 1},
			{0, 1, 0},
		}
		gameOfLifeAlt(board)
		fmt.Println("BOARD", board)
		for i := range board {
			for j := range board[i] {
				if board[i][j] != expected[i][j] {
					t.Errorf("Expected %d, got %d", expected[i][j], board[i][j])
				}
			}
		}
	})
}
