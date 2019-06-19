/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */
//  4 ms 96.08 %
func permute(nums []int) [][]int {
	// 排列问题
	rtn := make([][]int, 0)
	used := make([]bool, len(nums))
	dfs(nums, used, []int{}, &rtn)
	return rtn
}

func dfs(candidates []int, used []bool, curr []int, ans *[][]int) {
	// if s == len(candidates) {
	if len(curr) == len(candidates) {
		tmp := make([]int, len(candidates))
		copy(tmp, curr)
		(*ans) = append((*ans), tmp)
		return 
	}

	for i:=0; i<len(candidates); i++ {
		if used[i] == true { // 注意used的初始化，避免越界
			continue
		}
		used[i] = true
		curr = append(curr, candidates[i])
		dfs(candidates, used, curr, ans)
		curr = curr[:len(curr)-1]
		used[i] = false
	}
}

