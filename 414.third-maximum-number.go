/*
 * @lc app=leetcode id=414 lang=golang
 *
 * [414] Third Maximum Number
 */
//  4 ms 98.8 %
//  [2, 2, 3, 1]
// [1,2,2,5,3,5]
func thirdMax(nums []int) int {
	minInt := math.MinInt64 // -1 << 63
	max1,max2,max3 := minInt,minInt,minInt // int最小值
	for i:=0; i<len(nums); i++ {
		if nums[i] > max1 {
			max1,max2,max3 = nums[i],max1,max2
		}else if nums[i] != max1 && nums[i] > max2 { // 一定注意去重，防止少选出最大的3个值
			max2,max3 = nums[i],max2
		}else if nums[i] != max1 && nums[i] !=  max2 && nums[i] > max3 { 
			max3 = nums[i]
		}
	}
	if max3 == minInt {
		return max1
	}
	return max3
}

