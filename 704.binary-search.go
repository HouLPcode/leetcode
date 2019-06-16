/*
 * @lc app=leetcode id=704 lang=golang
 *
 * [704] Binary Search
 */
//  28 ms 99.68 %
func search(nums []int, target int) int {
    for i,j := 0,len(nums)-1; i<=j; {
		mid := (j-i)/2+i
		if nums[mid] == target {
			return mid
		}else if nums[mid] < target{
			i = mid + 1
		}else {
			j = mid - 1
		}
	}
	return -1
}

