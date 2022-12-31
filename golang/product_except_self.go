package leetcode

// 238. Product of Array Except Self

// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

// You must write an algorithm that runs in O(n) time and without using the division operation.

// Example 1:

// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]
// Example 2:

// Input: nums = [-1,1,0,-3,3]
// Output: [0,0,9,0,0]

// Constraints:

// 2 <= nums.length <= 105
// -30 <= nums[i] <= 30
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

// Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)

// This is solve the problem but it is not O(n) time complexity, it is O(n^2). And it would exceed time limit.
func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))

	for k1 := range nums {
		assigned := false
		for k2, v2 := range nums {
			if k1 == k2 {
				continue
			}

			if !assigned {
				answer[k1] = v2
				assigned = true
				continue
			}

			answer[k1] = answer[k1] * v2
		}
	}

	return answer
}

// This solves the problem, using O(n)
func productExceptSelfAlt(nums []int) []int {
	res := make([]int, len(nums))

	prefix := 1
	for i, v := range nums {
		res[i] = prefix
		prefix *= v
	}

	postfix := 1
	for i := len(res) - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}

	return res
}
