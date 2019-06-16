/*
 * @lc app=leetcode id=219 lang=golang
 *
 * [219] Contains Duplicate II
 */
//  12 ms 97.04 %
func containsNearbyDuplicate(nums []int, k int) bool {
	//想象成一个大小为k的滑动窗口，判断窗口内是否有重复元素
	// set存储k中数据，判重
	mset := make(map[int]bool, k)
	for i,v := range nums{
		if i > k { //窗口右移
			delete(mset, nums[i-k-1])
		}
		if mset[v] == true{
			return true
		} else {
			mset[v] = true
		}
	}
	return false
}

