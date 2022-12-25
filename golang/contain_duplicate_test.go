package leetcode

import "testing"

func TestContainDup(t *testing.T) {

	t.Log("Test: contain duplicate")
	t.Run("Example 1", func(t *testing.T) {
		s := []int{1, 2, 3, 1}
		expected := true
		got := containsDuplicate(s)
		if got != expected {
			t.Errorf("Expected %t, got %t", expected, got)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		s := []int{1, 2, 3, 4}
		expected := false
		got := containsDuplicate(s)
		if got != expected {
			t.Errorf("Expected %t, got %t", expected, got)
		}
	})

	t.Run("Example 3", func(t *testing.T) {
		s := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
		expected := true
		got := containsDuplicate(s)
		if got != expected {
			t.Errorf("Expected %t, got %t", expected, got)
		}
	})

	t.Run("Example 4", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		expected := false
		got := containsDuplicate(s)
		if got != expected {
			t.Errorf("Expected %t, got %t", expected, got)
		}
	})

	t.Run("Example 5", func(t *testing.T) {
		s := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
		expected := true
		got := containsDuplicate(s)
		if got != expected {
			t.Errorf("Expected %t, got %t", expected, got)
		}
	})

	t.Run("Example 6", func(t *testing.T) {
		s := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 2}
		expected := true
		got := containsDuplicate(s)
		if got != expected {
			t.Errorf("Expected %t, got %t", expected, got)
		}
	})

	t.Run("Example 7", func(t *testing.T) {
		s := []int{1}
		expected := false
		got := containsDuplicate(s)
		if got != expected {
			t.Errorf("Expected %t, got %t", expected, got)
		}
	})

	t.Run("Example 8", func(t *testing.T) {
		s := []int{1, 5, -2, -4, 0}
		expected := false
		got := containsDuplicate(s)
		if got != expected {
			t.Errorf("Expected %t, got %t", expected, got)
		}
	})

	t.Run("Example 9", func(t *testing.T) {
		s := []int{1000000000, 1000000000, 11}
		expected := true
		got := containsDuplicate(s)
		if got != expected {
			t.Errorf("Expected %t, got %t", expected, got)
		}
	})
}

func TestContainDupAlt(t *testing.T) {

	t.Log("Test: contain duplicate")
	t.Run("Example 1", func(t *testing.T) {
		s := []int{1, 2, 3, 1}
		expected := true
		got := containsDuplicateAlt1(s)
		if got != expected {
			t.Errorf("Expected %t, got %t", expected, got)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		s := []int{1, 2, 3, 4}
		expected := false
		got := containsDuplicateAlt1(s)
		if got != expected {
			t.Errorf("Expected %t, got %t", expected, got)
		}
	})

	t.Run("Example 3", func(t *testing.T) {
		s := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
		expected := true
		got := containsDuplicateAlt1(s)
		if got != expected {
			t.Errorf("Expected %t, got %t", expected, got)
		}
	})

	t.Run("Example 4", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		expected := false
		got := containsDuplicateAlt1(s)
		if got != expected {
			t.Errorf("Expected %t, got %t", expected, got)
		}
	})

	t.Run("Example 5", func(t *testing.T) {
		s := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
		expected := true
		got := containsDuplicateAlt1(s)
		if got != expected {
			t.Errorf("Expected %t, got %t", expected, got)
		}
	})

	t.Run("Example 6", func(t *testing.T) {
		s := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 2}
		expected := true
		got := containsDuplicateAlt1(s)
		if got != expected {
			t.Errorf("Expected %t, got %t", expected, got)
		}
	})

	t.Run("Example 7", func(t *testing.T) {
		s := []int{1}
		expected := false
		got := containsDuplicateAlt1(s)
		if got != expected {
			t.Errorf("Expected %t, got %t", expected, got)
		}
	})

	t.Run("Example 8", func(t *testing.T) {
		s := []int{1, 5, -2, -4, 0}
		expected := false
		got := containsDuplicateAlt1(s)
		if got != expected {
			t.Errorf("Expected %t, got %t", expected, got)
		}
	})

	t.Run("Example 9", func(t *testing.T) {
		s := []int{1000000000, 1000000000, 11}
		expected := true
		got := containsDuplicateAlt1(s)
		if got != expected {
			t.Errorf("Expected %t, got %t", expected, got)
		}
	})
}
