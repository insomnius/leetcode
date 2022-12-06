package leetcode

import "testing"

func TestLongestSubstringWithoutRepeatingChar(t *testing.T) {

	t.Log("Test: longest substring without repeating character")
	t.Run("Example 1", func(t *testing.T) {
		s := "abcabcbb"
		expected := 3
		got := lengthOfLongestSubstring(s)
		if got != expected {
			t.Errorf("Expected %d, got %d", expected, got)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		s := "bbbbb"
		expected := 1
		got := lengthOfLongestSubstring(s)
		if got != expected {
			t.Errorf("Expected %d, got %d", expected, got)
		}
	})

	t.Run("Example 3", func(t *testing.T) {
		s := "pwwkew"
		expected := 3
		got := lengthOfLongestSubstring(s)
		if got != expected {
			t.Errorf("Expected %d, got %d", expected, got)
		}
	})

}
