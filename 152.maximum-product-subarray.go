/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 */

 // 注意，统计三者中的最值时一定要有==，比如 2，2，1
func max(a,b,c int)int{
	if a >= b && a >= c{
		return a
	}
	if b >= a && b >= c{
		return b
	} 
	return c
}

func min(a,b,c int)int{
	if a <= b && a <= c{
		return a
	}
	if b <= a && b <= c{
		return b
	}
	return c
}

func maxProduct(nums []int) int {
	// 子序列必须连续，所以状态定义的时候必须包含当前值
	// 由于当前值包括正负情况，同时统计最大值和最小值

	// 状态函数: f(i) 表示包含当前值nums[i]的序列乘积的最大值或最小值，根据情况确定
	// 状态转移: f(i) = max(nums[i] , nums[i]*f(i-1) )
	// 最终结果 max( f(1),f(2),...   )，也可以在中间过程实时统计最大值
	// 该方法可以同时计算出最大乘积和最小乘积

	maxbuf, minbuf, res := nums[0], nums[0], nums[0]
	for _,v := range nums[1:]{//注意遍历的时候去除nums[0],已经遍历过了
		
		// 典型错误，计算minbuf时候maxbuf已经更新过了
		// maxbuf = max(v, maxbuf*v, minbuf*v)
		// minbuf = min(v, maxbuf*v, minbuf*v)
		
		maxbuf,minbuf = max(v, maxbuf*v, minbuf*v), min(v, maxbuf*v, minbuf*v)
		res = max(res,maxbuf,minbuf)
	}
	return res
}

