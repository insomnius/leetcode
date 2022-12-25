package leetcode

/*
217. Contains Duplicate

Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.



Example 1:

Input: nums = [1,2,3,1]
Output: true
Example 2:

Input: nums = [1,2,3,4]
Output: false
Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true


Constraints:

1 <= nums.length <= 105
-109 <= nums[i] <= 109
*/

// We use two pointer so it could iterate faster
func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	x, y := 0, (len(nums) - 1)
	keys := map[int]bool{}

	for x <= y {
		_, foundX := keys[nums[x]]
		_, foundY := keys[nums[y]]

		if foundX || foundY || (nums[x] == nums[y] && x != y) {
			return true
		}

		keys[nums[y]] = true
		keys[nums[x]] = true

		x++
		y--
	}

	return false
}

// We use normal iteration, which is works anyway
func containsDuplicateAlt1(nums []int) bool {
	keys := map[int]bool{}

	for i := 0; i < len(nums); i++ {
		_, found := keys[nums[i]]
		if found {
			return true
		}

		keys[nums[i]] = true
	}

	return false
}
