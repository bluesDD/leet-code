/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		comp := target - nums[i]
		if _, ok := numsMap[comp]; ok {
			answer := []int{i, numsMap[comp]}
			return answer
		}
	}
	return nil
}

// @lc code=end

