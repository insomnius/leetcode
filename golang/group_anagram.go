package leetcode

import "fmt"

// 49. Group Anagrams

// Given an array of strings strs, group the anagrams together. You can return the answer in any order.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

// Example 1:

// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
// Example 2:

// Input: strs = [""]
// Output: [[""]]
// Example 3:

// Input: strs = ["a"]
// Output: [["a"]]

// Constraints:

// 1 <= strs.length <= 104
// 0 <= strs[i].length <= 100
// strs[i] consists of lowercase English letters.

// This solution actually got time limited, and it's not the best solution.
func groupAnagrams(strs []string) [][]string {
	anagrams := [][]string{}
	if len(strs) == 1 {
		return append(anagrams, strs)
	}

	for _, str := range strs {
		found := false
		for k, anagram := range anagrams {
			if isAnagram(str, anagram[0]) {
				anagrams[k] = append(anagrams[k], str)
				found = true
				break
			}
		}

		if found {
			continue
		}

		anagrams = append(anagrams, []string{str})
	}

	return anagrams
}

func groupAnagramsAlt(strs []string) [][]string {
	anagrams := [][]string{}
	if len(strs) == 1 {
		return append(anagrams, strs)
	}

	anagramHashMap := map[string][]string{}

	for _, str := range strs {
		arr := make([]int, 26)
		for _, char := range str {
			arr[char-'a']++
		}

		arrToS := fmt.Sprintf("%v", arr)
		anagramHashMap[arrToS] = append(anagramHashMap[arrToS], str)
	}

	for _, hashMap := range anagramHashMap {
		anagramArr := []string{}

		for _, v := range hashMap {
			anagramArr = append(anagramArr, v)
		}
		anagrams = append(anagrams, anagramArr)
	}

	return anagrams
}
