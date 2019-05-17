/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */
 //测试集 [5,1,3]\n5
//       [1,3]\n3
//       [5,1,2,3,4]\n1
// 分析思路：
// 1.找到有序的一侧nums[l]<=nums[mid]肯定左边有序，nums[mid]<=nums[r]肯定右边有序
// 2.target和区间边界值比较，决定在不在该区间
func search(nums []int, target int) int {
	//直接二分查找，对mid的转移做变化
	l,r := 0, len(nums)-1
	for l <= r{
		mid := (r-l)/2 + l
		if nums[mid] == target{
			return mid
		}else if (nums[l] <= nums[mid] && nums[l] <= target && target < nums[mid]) || //这里列出肯定落在左区间的条件，剩余其他情况肯定落在右区间
			(nums[mid] <= nums[r] && 
				((target > nums[mid] && target > nums[r])||
				(target < nums[mid] && target < nums[r]))){ //注意target可能等于left或right
			r = mid - 1	
		}else {
			l = mid + 1
		}
	}
	return -1
}

