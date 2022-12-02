package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidParentheses(t *testing.T) {

	t.Log("Test: valid parentheses")
	t.Run("Example 1", func(t *testing.T) {
		s := "()"
		result := isValid(s)
		assert.Equal(t, true, result)
	})

	t.Run("Example 2", func(t *testing.T) {
		s := "{}"
		result := isValid(s)
		assert.Equal(t, true, result)
	})

	t.Run("Example 3", func(t *testing.T) {
		s := "[]"
		result := isValid(s)
		assert.Equal(t, true, result)
	})

	t.Run("Example 4", func(t *testing.T) {
		s := "[](){}{}()[]"
		result := isValid(s)
		assert.Equal(t, true, result)
	})

	t.Run("Example 5", func(t *testing.T) {
		s := "[]()()}[]}()[]"
		result := isValid(s)
		assert.Equal(t, false, result)
	})

	t.Run("Example 6", func(t *testing.T) {
		s := "([)]"
		result := isValid(s)
		assert.Equal(t, false, result)
	})

	t.Run("Example 7", func(t *testing.T) {
		s := "{[]}"
		result := isValid(s)
		assert.Equal(t, true, result)
	})

	t.Run("Example 8", func(t *testing.T) {
		s := "({()})"
		result := isValid(s)
		assert.Equal(t, true, result)
	})
}
