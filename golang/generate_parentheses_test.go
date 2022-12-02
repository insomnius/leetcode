package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParentheses(t *testing.T) {
	t.Log("Test: generate parentheses")
	t.Run("Example 1", func(t *testing.T) {
		n := 3
		result := generateParenthesis(n)
		assert.Equal(t, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}, result)
	})

	t.Run("Example 2", func(t *testing.T) {
		n := 1
		result := generateParenthesis(n)
		assert.Equal(t, []string{"()"}, result)
	})
}
