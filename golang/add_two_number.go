package leetcode

// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
// Example 1:
// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.
//
// Example 2:
// Input: l1 = [0], l2 = [0]
// Output: [0]
//
// Example 3:
// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// Output: [8,9,9,9,0,0,0,1]
//
// Constraints:
// The number of nodes in each linked list is in the range [1, 100].
// 0 <= Node.val <= 9
// It is guaranteed that the list represents a number that does not have leading zeros.
// https://leetcode.com/problems/add-two-numbers/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummyHead = new(ListNode)
	curr := dummyHead
	var carry int
	for l1 != nil || l2 != nil || carry != 0 {
		l1Val := 0
		if l1 != nil {
			l1Val = l1.Val
		}
		l2Val := 0
		if l2 != nil {
			l2Val = l2.Val
		}
		columnSum := l1Val + l2Val + carry
		carry = columnSum / 10
		newNode := &ListNode{Val: columnSum % 10, Next: nil}
		curr.Next = newNode
		curr = newNode
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return dummyHead.Next
}
