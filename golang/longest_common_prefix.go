package leetcode

// 14. Longest Common Prefix

// Write a function to find the longest common prefix string amongst an array of strings.

// If there is no common prefix, return an empty string "".

// Example 1:

// Input: strs = ["flower","flow","flight"]
// Output: "fl"
// Example 2:

// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

// Constraints:

// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] consists of only lowercase English letters.

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	compiledStr := ""
	i := 0
	firstS := strs[0]
	loop := true

	for loop && i < len(firstS) {
		for _, s := range strs[1:] {
			if i >= len(s) {
				loop = false
				break
			}

			if string(s[i]) != string(firstS[i]) {
				loop = false
				break
			}
		}

		if loop {
			compiledStr += string(firstS[i])
		}
		i++
	}

	return compiledStr
}

// Shorter and much much better
func longestCommonPrefixAlt(strs []string) string {
	compiledStr := ""
	i := 0

	for i < len(strs[0]) {
		for _, s := range strs {
			if i == len(s) || s[i] != strs[0][i] {
				return compiledStr
			}
		}
		compiledStr += string(strs[0][i])
		i++
	}

	return compiledStr
}
