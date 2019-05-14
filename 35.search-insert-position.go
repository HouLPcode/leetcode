/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */
func searchInsert(nums []int, target int) int {
	//首先二分查找是否有target
	i,j := 0, len(nums)-1
	for i <= j{
		mid := (j-i)/2+i
		if nums[mid] == target{
			return mid
		}else if nums[mid] < target{
			i = mid + 1
		}else{
			j = mid -1
		}
	}
	//跳出循环表示没有找到, i>j
	return i // 要插入的位置是较大的那个位置
}

