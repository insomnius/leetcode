package leetcode

/*
55. Jump Game

You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.



Example 1:

Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
Example 2:

Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.


Constraints:

1 <= nums.length <= 104
0 <= nums[i] <= 105

*/

func canJump(nums []int) bool {
	index := 0
	length := len(nums) - 1

	for index < length {
		i := nums[index]
		addIndex := 0
		max := 0

		if index+i >= length {
			return true
		}

		for i > 0 {
			if index+i+nums[index+i] > max {
				max = index + i + nums[index+i]
				addIndex = i
			}
			i--
		}

		if addIndex == 0 {
			return false
		}

		index += addIndex
	}

	return true
}

func canJumpalt(nums []int) bool {
	index := 0

	for index < len(nums)-1 {
		i := nums[index]
		addIndex := 0
		max := 0

		if index+i >= len(nums)-1 {
			return true
		}

		for i > 0 {
			if index+i+nums[index+i] > max {
				max = index + i + nums[index+i]
				addIndex = i
			}
			i--
		}

		if addIndex == 0 {
			return false
		}

		index += addIndex
	}

	return true
}

// GREEDY ALGORITHM IS A KEY
func canJumpaltGreedy(nums []int) bool {
	maxLocation := 0

	for i := 0; i < len(nums); i++ {
		if i > maxLocation {
			return false
		}

		if i+nums[i] > maxLocation {
			maxLocation = i + nums[i]
		}
	}
	return true
}
