package leetcode

import "testing"

func TestTopKFrequen(t *testing.T) {

	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	result := topKFrequent(nums, k)

	if len(result) != k {
		t.Errorf("Expected %d, got %d", k, len(result))
	}

}
