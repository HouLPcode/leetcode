/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */
func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	generate(0, []int{}, candidates, target, 0, &result)
	return result
}

// 第k个数
func generate(k int, path []int, candidates []int, target int, sum int, result *[][]int){
	if k == len(candidates){
		return 
	}
	pathCopy := make([]int, len(path))
	copy(pathCopy, path)
	
	if sum == target{
		*result = append((*result), pathCopy)
		return 
	}else if sum > target{
		return
	}
	// 使用该元素，单次或者多次
	for ; k<len(candidates); k++ {
		generate(k, append(path,candidates[k]), candidates, target, sum+candidates[k], result)
		copy(path, pathCopy)
	} 
	// 不使用该元素
	for ; k<len(candidates); k++ {
		generate(k, path, candidates, target, sum, result)
	} 
}
