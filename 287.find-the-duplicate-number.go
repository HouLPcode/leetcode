/*
 * @lc app=leetcode id=287 lang=golang
 *
 * [287] Find the Duplicate Number
 */
//  4 ms  99.6 %
//  环的入口
func findDuplicate(nums []int) int {
	// n个数填充n+1个元素，肯定有环
	fast := nums[nums[0]]
	slow := nums[0]
	for slow != fast {
		fast = nums[nums[fast]]
		slow = nums[slow]
	}
	slow = nums[0]
	fast = nums[fast]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
