package leetcode

import (
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {

	t.Log("Test: longest valid parentheses")
	t.Run("Example 1", func(t *testing.T) {
		s := "(()"
		result := longestValidParentheses(s)
		if result != 2 {
			t.Errorf("Expected 2, but it was %d instead.", result)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		s := ")()())"
		result := longestValidParentheses(s)
		if result != 4 {
			t.Errorf("Expected 4, but it was %d instead.", result)
		}
	})

	t.Run("Example 3", func(t *testing.T) {
		s := ""
		result := longestValidParentheses(s)
		if result != 0 {
			t.Errorf("Expected 0, but it was %d instead.", result)
		}
	})

	t.Run("Example 4", func(t *testing.T) {
		s := "(((())))))"
		result := longestValidParentheses(s)
		if result != 8 {
			t.Errorf("Expected 8, but it was %d instead.", result)
		}
	})

	t.Run("Example 5", func(t *testing.T) {
		s := "()(()"
		result := longestValidParentheses(s)
		if result != 2 {
			t.Errorf("Expected 2, but it was %d instead.", result)
		}
	})

	t.Run("Example 6", func(t *testing.T) {
		s := "()()()()()()()()())(((((())))))"
		result := longestValidParentheses(s)
		if result != 18 {
			t.Errorf("Expected 18, but it was %d instead.", result)
		}
	})

	t.Run("Example 7", func(t *testing.T) {
		s := "()(((()"
		result := longestValidParentheses(s)
		if result != 2 {
			t.Errorf("Expected 2, but it was %d instead.", result)
		}
	})

	t.Run("Example 8", func(t *testing.T) {
		s := "(()()"
		result := longestValidParentheses(s)
		if result != 4 {
			t.Errorf("Expected 4, but it was %d instead.", result)
		}
	})
}
