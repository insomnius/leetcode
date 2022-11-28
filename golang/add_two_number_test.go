package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumber(t *testing.T) {
	t.Log("Test: add two number")
	t.Run("Example 1", func(t *testing.T) {
		l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
		l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
		result := addTwoNumbers(l1, l2)
		assert.Equal(t, []int{7, 0, 8}, listNodeToArray(result))
	})
}

func listNodeToArray(listNode *ListNode) []int {
	var result []int
	for listNode != nil {
		result = append(result, listNode.Val)
		listNode = listNode.Next
	}

	return result
}
