package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJumpGame(t *testing.T) {
	t.Log("Test: jump game ")

	t.Run("Example 1", func(t *testing.T) {
		nums := []int{2, 3, 1, 1, 4}
		result := canJump(nums)
		assert.Equal(t, true, result)
	})

	t.Run("Example 2", func(t *testing.T) {
		nums := []int{3, 2, 1, 0, 4}
		result := canJump(nums)
		assert.Equal(t, false, result)
	})

	t.Run("Example 3", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		result := canJump(nums)
		assert.Equal(t, true, result)
	})
}

func TestJumpGameAlt(t *testing.T) {
	t.Log("Test: jump game ")

	t.Run("Example 1", func(t *testing.T) {
		nums := []int{2, 3, 1, 1, 4}
		result := canJumpalt(nums)
		assert.Equal(t, true, result)
	})

	t.Run("Example 2", func(t *testing.T) {
		nums := []int{3, 2, 1, 0, 4}
		result := canJumpalt(nums)
		assert.Equal(t, false, result)
	})

	t.Run("Example 3", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		result := canJumpalt(nums)
		assert.Equal(t, true, result)
	})
}
