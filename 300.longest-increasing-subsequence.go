/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */
func lengthOfLIS(nums []int) int {
	// 注意是非连续的子序列，转移方程的设计
	// 状态方程 f(i) 包含第i元素的长度
	// 转移方程 f(i）= max(f(j) + 1) ,
	          // 其中 j= 0...i-1 且 nums[j] < nums[i]
	if nums==nil || len(nums) == 0{
		return 0
	}
	lenMax := 1
	lis := make([]int,len(nums),len(nums))//申请数组时不能使用变量长度？？？
	for i:=0; i< len(nums); i++{
		curMax := 1
		for j:=0;j<i;j++{
			if nums[j] < nums[i]{
				curMax = max(curMax,lis[j]+1)
			}
		}
		lis[i] = curMax
		if curMax > lenMax{
			lenMax = curMax
		}
	}
	return lenMax
}

func max(a,b int)int{
	if a > b{
		return a
	}
	return b
}

