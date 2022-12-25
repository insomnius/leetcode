package leetcode

// 242. Valid Anagram

// Given two strings s and t, return true if t is an anagram of s, and false otherwise.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

// Example 1:

// Input: s = "anagram", t = "nagaram"
// Output: true
// Example 2:

// Input: s = "rat", t = "car"
// Output: false

// Constraints:

// 1 <= s.length, t.length <= 5 * 104
// s and t consist of lowercase English letters.

// Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var x, y int
	charX := map[byte]int{}
	charY := map[byte]int{}

	for x < len(s) {
		charX[s[x]]++
		charY[t[y]]++
		x++
		y++
	}

	for kX, vX := range charX {
		vY, found := charY[kX]
		if !found {
			return false
		}

		if vX != vY {
			return false
		}
	}

	return true
}
