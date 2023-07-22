package leetcode

import "sort"

/*
128. Longest Consecutive Sequence
Medium
17.2K
743
Companies
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.



Example 1:

Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
Example 2:

Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9


Constraints:

0 <= nums.length <= 105
-109 <= nums[i] <= 109

*/

func longestConsecutive(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	if len(nums) < 2 {
		return 1
	}

	sort.Ints(nums)

	max := 0
	tempTotal := 0

	for i := 0; i < len(nums); i++ {
		tempTotal++

		if i == 0 {
			continue
		}

		if nums[i] == nums[i-1] {
			tempTotal--
			continue
		}

		if !(nums[i]-nums[i-1] == 1) {
			tempTotal = 1
		}

		if tempTotal > max {
			max = tempTotal
		}
	}

	if tempTotal > max {
		max = tempTotal
	}

	return max
}
