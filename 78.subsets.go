/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */
 // 法1，回溯法
 // 法2，位运算，合理的值表示所有情况
func subsets(nums []int) [][]int {
	// 回溯法思路：
	// generate i 对i元素进行放入和不放入的操作，后递归调动i+1 ，
	// 结束条件：i大于最后一个元素
	tmp := []int{}
	result := [][]int{}
	generate(0, tmp, nums, &result)
	return result
}

func generate(i int, tmp, nums []int, result *[][]int){
	if i == len(nums){
		return 
	}
	// 不放入i元素
	*result = append((*result), tmp)
	generate(i+1, tmp, nums, result)
	// 放入i元素
	tmp = append(tmp, nums[i])
	*result = append((*result), tmp)	
	generate(i+1, tmp, nums, result)
}

