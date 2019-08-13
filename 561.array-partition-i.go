import "sort"

/*
 * @lc app=leetcode id=561 lang=golang
 *
 * [561] Array Partition I
 */
//  (72 ms) √ Your runtime beats 76.47 %
// 先排序，后计算 nums[0,2,4,6...]的和
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}
