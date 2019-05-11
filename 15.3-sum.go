/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */
 func threeSum(nums []int) [][]int {
	rtn := [][]int{}
	if len(nums) < 3 { // nums == nil也可以用len
		return rtn
	}
	// 本题难点是怎么保证输出结果不重复 [-1,0,1]和[-1,1,0]表示的是同一个
	// ****输出结果不能冗余

	sort.Ints(nums) //库函数，默认从小到大
	//go中slice的index必须非负
	for a := 0; a < len(nums)-2; a++ {
		if a > 0 && nums[a] == nums[a-1] { // 结果中去重,注意a>0条件
			continue
		}
		for b, c := a+1, len(nums)-1; b < c; {
			if nums[a]+nums[b]+nums[c] == 0 { // 找到结果
				rtn = append(rtn, []int{nums[a], nums[b], nums[c]})
				// 找到后，后续可能还有与a匹配的结果，继续寻找
				for b++; nums[b] == nums[b-1] && b < c; b++ {
				} //无限的++可能导致nums越界
				for c--; nums[c] == nums[c+1] && b < c; c-- {
				}
			} else if nums[a]+nums[b]+nums[c] < 0 { // 注意一定是else if， 否则在b或c有重复值得时候会死循环 [-1,0,0,1]
				b++
			} else if nums[a]+nums[b]+nums[c] > 0 {
				c--
			}
		}
	}
	return rtn
}

// 0,-1,1
// 0,0,-1,-1,1,1
// 0,0,-1,-1,1,1,-2,-2,2,4


