package leetcode

/*
Given a string s, find the length of the longest
substring
 without repeating characters.

Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.


Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.

*/

func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}

	left := 0
	right := 0
	max := 0

	var stack []string

	checkStack := func(str string) bool {
		for _, s := range stack {
			if s == str {
				return true
			}
		}

		return false
	}

	for right < len(s) {
		if checkStack(string(s[right])) {
			if left < right {
				left++
			}
			stack = stack[1:]
			continue
		}
		stack = append(stack, string(s[right]))

		if len(stack) > max {
			max = len(stack)
		}

		right++
	}

	return max
}
