import "sort"

/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */
//   8 ms, faster than 44.44%
// [0,2,1,-3]\n1
func threeSumClosest(nums []int, target int) int {
	// 排序后遍历每个元素，再剩余元素中，首尾指针计算三个元素的和
	sort.Ints(nums)
	rtn := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		for l, r := i+1, len(nums)-1; l < r; {
			sum := nums[i] + nums[l] + nums[r]
			if subAbs(rtn, target) > subAbs(sum, target) { // rtn初始为0的影响？？？
				rtn = sum
			}
			if sum < target {
				l++
			} else if sum > target {
				r--
			} else {
				return target // 等于的时候最小，直接返回
			}
		}
	}
	return rtn
}

func subAbs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
