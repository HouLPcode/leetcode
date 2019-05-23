/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */
 // 法1，回溯法
 // 法2，位运算，合理的值表示所有情况
 // [9,0,3,5,7] 回溯法的顺序不满足答案要求？？？
func subsets(nums []int) [][]int {
	// 回溯法思路：
	// generate i 对i元素进行放入和不放入的操作，后递归调动i+1 ，
	// 结束条件：i大于最后一个元素
	tmp := []int{}
	result := [][]int{[]int{}}//默认含有一个空子集
	sort.Ints(nums)
	generate(0, tmp, nums, &result)
	return result
}

func generate(i int, tmp, nums []int, result *[][]int){
	if i == len(nums){
		return 
	}

	// 不放入i元素
	// *result = append((*result), tmp)
	generate(i+1, tmp, nums, result)
	
	// 放入i元素
	tmp = append(tmp, nums[i])
	// 注意，只有带i元素的结果需要，不带i元素的情况在上层递归前已经插入了
	// 注意，少了一个元素也没有的子集
	data := make([]int, len(tmp))
	copy(data, tmp)
	*result = append((*result), data)	
	generate(i+1, tmp, nums, result)
}

