/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 */
//  (804 ms) √ Your runtime beats 80.56 %
//  []\n0
// [1,-1]\n1
// 最大值队列
//  入队：空，直接放入。值>=队尾，弹出队尾，继续比较。值<队尾，放入
func maxSlidingWindow(nums []int, k int) []int {
	rtn := make([]int, 0, len(nums))
	if len(nums) == 0 || k == 0 {
		return rtn
	}
	maxQueue := make([]int, 0, k)
	maxQueue = append(maxQueue, 0)
	curMax := nums[0]
	rtn = append(rtn, curMax)
	for i := 1; i < len(nums); i++ {
		if len(maxQueue) > 0 && i == maxQueue[0]+k {
			// 出队
			maxQueue = maxQueue[1:]
		}
		// 入队
		for len(maxQueue) > 0 && nums[i] >= nums[maxQueue[len(maxQueue)-1]] { // 值>=队尾,出来
			maxQueue = maxQueue[:len(maxQueue)-1]
		}
		maxQueue = append(maxQueue, i) // 注意队列里面放的是下标，不是内容。方便出栈时的比较
		curMax = nums[maxQueue[0]]
		rtn = append(rtn, curMax)
	}
	return rtn[k-1:] // 舍头
}
