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
	
	if sum == target{
		pathCopy := make([]int, len(path))
		copy(pathCopy, path)
		*result = append((*result), pathCopy)
		return 
	}else if sum > target{
		return
	}

	// 使用该元素，此处不是k+1，表示可以使用多次该元素
	path = append(path,candidates[k])
	generate(k, path, candidates, target, sum+candidates[k], result)
	// 不使用该元素
	generate(k+1, path[:len(path)-1], candidates, target, sum, result)
}
