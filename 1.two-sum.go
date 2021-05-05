/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	for idx, num := range nums {
		for j := idx + 1; j < len(nums); j++ {
			if num+nums[j] == target {
				return []int{idx, j}
			}
		}
	}
	return nil
}

// @lc code=end

