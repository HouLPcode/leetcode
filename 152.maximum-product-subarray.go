/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 */

// 典型错误

//  ✘ testcase: '[-2,3,-4]'
//   ✘ answer: 3
//   ✘ expected_answer: 24
func maxProduct(nums []int) int {
	// 子序列必须连续，怎么决定子序列，必须包含当前值？？？？
	// 需要同时统计最大值和最小值？？？

	// 第一次尝试：
	// 状态函数: f(i) 表示包含最后一个值nums[i]的序列的乘积最大值
	// 状态转移: f(i) = max(nums[i] , nums[i]*f(i-1) )
	// 最终结果 max( f(1),f(2),...   )
	// 这是一种典型错误，贪心思想不能解决该问题

	rnt := make(map[int]int)
	for i :=0; i < len(nums); i++ {// 不能直接range 数字
		help(i,nums,rnt)
	}

	max := nums[0]
	for _,v := range rnt{
		if max < v{
			max = v
		}
	}
	return max
}

// 递归 + 缓存
// ???缓存必须用map？？？用slice怎么确认是零值还是计算值
func help(i int, nums []int, buf map[int]int) int{
	if i == 0 {
		buf[i] = nums[i]
		return nums[i]
	}
	if v, ok := buf[i]; ok{
		return v
	}
	p := help(i-1, nums, buf) * nums[i]
	if nums[i] > p{
		buf[i] = nums[i]
		return nums[i]
	}
	buf[i] = p
	return p
}

