package leetcode

import "testing"

func TestGroupAnagram(t *testing.T) {

	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(strs)

	if len(result) != 3 {
		t.Errorf("Expected 3, got %d", len(result))
	}

	if len(result[0]) != 3 {
		t.Errorf("Expected 3, got %d", len(result[0]))
	}

	if len(result[1]) != 2 {
		t.Errorf("Expected 2, got %d", len(result[1]))
	}

	if len(result[2]) != 1 {
		t.Errorf("Expected 1, got %d", len(result[2]))
	}

}
