/*
 * @lc app=leetcode id=724 lang=golang
 *
 * [724] Find Pivot Index
 */
//  20 ms 94.51 %
func pivotIndex(nums []int) int {
	// 遍历一边得到总数
	// 第二遍找index
	sum := 0
	for _,v := range nums{
		sum += v
	}
	
	suml := 0
	sumr := sum //初始值时。左为0，右为总和，在左右比较前减去当前值
	for index := 0; index<len(nums); index++{
		sumr -= nums[index] //注意右侧总和延后计算，
		if suml == sumr {
			return index
		}else {
			suml += nums[index]
		}
	}
	return -1
}

