/*
 * @lc app=leetcode id=494 lang=golang
 *
 * [494] Target Sum
 */
// 参考leetcode 113
// 测试集 [1,0] 1
// DFS 递归思路
func findTargetSumWays(nums []int, S int) int {
	//构成二叉树，得到所有路径的sum，然后遍历sum值为S的个数
	cnt := 0
	find(nums, S, &cnt)
	return cnt
}

func find(nums []int, S int, cnt *int){
	if len(nums) == 0{
		return
	}
	//此处条件必须是数组最后一个数字且为所求值
	if len(nums)==1 && nums[0] == S{//此处需要考虑+-S,因为find传参数的时候没有计算当前值得正负，传的是上一个函数的正负
		// 最后一个元素达到想要的值，计数器++ 
		(*cnt)++
	}
	if len(nums) == 1 && nums[0] == -S{
		(*cnt)++
	}

	find(nums[1:], S-nums[0], cnt)
	find(nums[1:], S+nums[0], cnt)
}

