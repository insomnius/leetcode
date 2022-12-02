package leetcode

/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

Example 1:

Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
Example 2:

Input: n = 1
Output: ["()"]


Constraints:

1 <= n <= 8
*/

func generateParenthesis(n int) []string {
	return recParenthesis(n, "", 0, 0)
}

func recParenthesis(n int, parentheses string, left, right int) []string {
	var parenthesesList []string

	// if len(parentheses) == n*2 {
	// 	return []string{parentheses}
	// } else if left < n && left < right {
	// 	parenthesesList = append(parenthesesList, recParenthesis(n, "("+parentheses, left+1, right)...)
	// } else if right < n && right < left {
	// 	parenthesesList = append(parenthesesList, recParenthesis(n, parentheses+")", left, right+1)...)
	// 	if left < n {
	// 		fmt.Println("(" + parentheses)
	// 		parenthesesList = append(parenthesesList, recParenthesis(n, "("+parentheses, left+1, right)...)
	// 	}
	// }

	if len(parentheses) == n*2 {
		return []string{parentheses}
	}

	if left < n {
		parenthesesList = append(parenthesesList, recParenthesis(n, parentheses+"(", left+1, right)...)
	}

	if right < left {
		parenthesesList = append(parenthesesList, recParenthesis(n, parentheses+")", left, right+1)...)
	}

	return parenthesesList
}
