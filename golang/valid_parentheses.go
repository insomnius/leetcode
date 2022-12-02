package leetcode

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
*/

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	var arr []string

	for _, r := range s {
		switch symbol := string(r); symbol {
		case "(":
			arr = append(arr, "(")
		case ")":
			if len(arr) == 0 || arr[len(arr)-1] != "(" {
				return false
			}
			arr = arr[:len(arr)-1]
		case "{":
			arr = append(arr, "{")
		case "}":
			if len(arr) == 0 || arr[len(arr)-1] != "{" {
				return false
			}
			arr = arr[:len(arr)-1]
		case "[":
			arr = append(arr, "[")
		case "]":
			if len(arr) == 0 || arr[len(arr)-1] != "[" {
				return false
			}
			arr = arr[:len(arr)-1]
		}
	}

	return len(arr) == 0
}
