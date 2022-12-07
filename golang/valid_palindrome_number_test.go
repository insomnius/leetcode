package leetcode

import "testing"

func TestValidPalindrome(t *testing.T) {
	t.Log("Test: valid palindrome number")
	t.Run("Example 1", func(t *testing.T) {
		s := 101
		expected := true
		got := isPalindrome(s)
		if got != expected {
			t.Errorf("Expected %t, got %t", expected, got)
		}
	})

	t.Run("Example 1", func(t *testing.T) {
		s := 1001
		expected := true
		got := isPalindrome(s)
		if got != expected {
			t.Errorf("Expected %t, got %t", expected, got)
		}
	})
}
