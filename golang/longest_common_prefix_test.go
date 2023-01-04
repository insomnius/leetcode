package leetcode

import "testing"

func TestLongestCommonPrefix(t *testing.T) {

	t.Run("Example 1", func(t *testing.T) {
		strs := []string{"flower", "flow", "flight"}
		expected := "fl"
		got := longestCommonPrefix(strs)
		if got != expected {
			t.Errorf("Expected %s, got %s", expected, got)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		strs := []string{"dog", "racecar", "car"}
		expected := ""
		got := longestCommonPrefix(strs)
		if got != expected {
			t.Errorf("Expected %s, got %s", expected, got)
		}
	})
}

func TestLongestCommonPrefixAlt(t *testing.T) {

	t.Run("Example 1", func(t *testing.T) {
		strs := []string{"flower", "flow", "flight"}
		expected := "fl"
		got := longestCommonPrefixAlt(strs)
		if got != expected {
			t.Errorf("Expected %s, got %s", expected, got)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		strs := []string{"dog", "racecar", "car"}
		expected := ""
		got := longestCommonPrefixAlt(strs)
		if got != expected {
			t.Errorf("Expected %s, got %s", expected, got)
		}
	})
}
