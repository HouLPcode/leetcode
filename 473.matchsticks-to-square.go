/*
 * @lc app=leetcode id=473 lang=golang
 *
 * [473] Matchsticks to Square
 */
 func makesquare(nums []int) bool {
	//可以理解成组合问题，回溯法遍历所有的组合可能，剪枝减少不必要的检查
	if len(nums) < 4{
		return false
	}
	sum := 0
	for i:=0; i<len(nums); i++ {
		sum += nums[i]
	}
	if sum % 4 != 0{
		return false
	}
	//nums 从大到小排序,能够剪枝，减少不必要的回溯，有点贪心的意思
	sort.Ints(nums)
	bucket := make([]int, 4, 4)
	return generate(len(nums)-1, nums, sum/4, bucket)
}
// i表示nums中的第几个元素，target是每边的长度，bucket为四个边目前的长度
func generate(i int, nums []int, target int, bucket []int) bool{
	if i == -1{// 注意，此处是len的时候比较，不是len-1，nums[len-1]的元素需要放入bucket中之后才能比较
		return bucket[0] == target &&
			bucket[1] == target &&
			bucket[2] == target &&
			bucket[3] == target
	}
	//回溯，遍历所有可能
	for k,val := range bucket{
		if val + nums[i] > target{
			//return false 此处不应直接返回false，应该将nums[i]放入其他bucket继续
			continue
		}
		bucket[k] += nums[i]
		// 递归的时候不用管i范围
		// 此处是true的时候找到合适的组合，直接返回即可，不需要后续的遍历了
		if generate(i-1, nums, target, bucket) == true{
			return true//已经找到合适的组合，直接跳出递归，返回true
		}
		bucket[k] -= nums[i] // 尝试完后恢复到原先状态
	}
	return false
}
