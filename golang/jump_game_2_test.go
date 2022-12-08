package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJumpGame2(t *testing.T) {
	t.Log("Test: jump game 2")
	t.Run("Example 1", func(t *testing.T) {
		nums := []int{2, 3, 1, 1, 4}
		result := jump2(nums)
		assert.Equal(t, 2, result)
	})

	t.Run("Example 2", func(t *testing.T) {
		nums := []int{2, 3, 0, 1, 4}
		result := jump2(nums)
		assert.Equal(t, 2, result)
	})

	t.Run("Example 3", func(t *testing.T) {
		nums := []int{1, 2, 3, 2, 3, 4, 6, 7}
		result := jump2(nums)
		assert.Equal(t, 4, result)
	})

	t.Run("Example 4", func(t *testing.T) {
		nums := []int{2, 3, 1}
		result := jump2(nums)
		assert.Equal(t, 1, result)
	})

	t.Run("Example 5", func(t *testing.T) {
		nums := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 0}
		result := jump2(nums)
		assert.Equal(t, 2, result)
	})

	t.Run("Example 6", func(t *testing.T) {
		nums := []int{3, 4, 3, 2, 5, 4, 3}
		result := jump2(nums)
		assert.Equal(t, 3, result)
	})

	t.Run("Example 7", func(t *testing.T) {
		nums := []int{9, 8, 2, 2, 0, 2, 2, 0, 4, 1, 5, 7, 9, 6, 6, 0, 6, 5, 0, 5}
		result := jump2(nums)
		assert.Equal(t, 3, result)
	})

}

func TestJumpGame2Alt(t *testing.T) {
	t.Log("Test: jump game 2")
	t.Run("Example 1", func(t *testing.T) {
		nums := []int{2, 3, 1, 1, 4}
		result := jump2alt1(nums)
		assert.Equal(t, 2, result)
	})

	t.Run("Example 2", func(t *testing.T) {
		nums := []int{2, 3, 0, 1, 4}
		result := jump2alt1(nums)
		assert.Equal(t, 2, result)
	})

	t.Run("Example 3", func(t *testing.T) {
		nums := []int{1, 2, 3, 2, 3, 4, 6, 7}
		result := jump2alt1(nums)
		assert.Equal(t, 4, result)
	})

	t.Run("Example 4", func(t *testing.T) {
		nums := []int{2, 3, 1}
		result := jump2alt1(nums)
		assert.Equal(t, 1, result)
	})

	t.Run("Example 5", func(t *testing.T) {
		nums := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 0}
		result := jump2alt1(nums)
		assert.Equal(t, 2, result)
	})

	t.Run("Example 6", func(t *testing.T) {
		nums := []int{3, 4, 3, 2, 5, 4, 3}
		result := jump2alt1(nums)
		assert.Equal(t, 3, result)
	})

	t.Run("Example 7", func(t *testing.T) {
		nums := []int{9, 8, 2, 2, 0, 2, 2, 0, 4, 1, 5, 7, 9, 6, 6, 0, 6, 5, 0, 5}
		result := jump2alt1(nums)
		assert.Equal(t, 3, result)
	})

}
