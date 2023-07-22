package leetcode

import "fmt"

// NOTES: Always using stack and index to solve this problem. Since we need to find max, then the stack should be empty when we find the max. So we need to push -1 to the stack first.

/*
Given a string containing just the characters '(' and ')', return the length of the longest valid (well-formed) parentheses
substring
.



Example 1:

Input: s = "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()".
Example 2:

Input: s = ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()".
Example 3:

Input: s = ""
Output: 0


Constraints:

0 <= s.length <= 3 * 104
s[i] is '(', or ')'.
*/

func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}

	stack := []int{-1}
	max := 0

	for i := 0; i < len(s); i++ {
		fmt.Println(stack)
		if string(s[i]) == "(" {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				if max < i-stack[len(stack)-1] {
					max = i - stack[len(stack)-1]
				}
			}
		}
	}

	return max
}
