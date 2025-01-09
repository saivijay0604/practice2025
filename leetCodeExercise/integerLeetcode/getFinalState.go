package leetcode

// You are given an integer array nums, an integer k, and an integer multiplier.
// You need to perform k operations on nums. In each operation:
// Find the minimum value x in nums. If there are multiple occurrences of the minimum value, select the one that appears first.
// Replace the selected minimum value x with x * multiplier.

// Input: nums = [2,1,3,5,6], k = 5, multiplier = 2

// Output: [8,4,6,5,6]

// Explanation:

// Operation	Result
// After operation 1	[2, 2, 3, 5, 6]
// After operation 2	[4, 2, 3, 5, 6]
// After operation 3	[4, 4, 3, 5, 6]
// After operation 4	[4, 4, 6, 5, 6]
// After operation 5	[8, 4, 6, 5, 6]

func GetFinalState(nums []int, k int, multiplier int) []int {
	if len(nums) == 0 {
		return nil
	}

	for i := 0; i < k; i++ {
		minIndex := 0
		for j := 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j

			}
		}
		nums[minIndex] *= multiplier
	}
	return nums
}
