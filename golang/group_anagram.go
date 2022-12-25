package leetcode

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
