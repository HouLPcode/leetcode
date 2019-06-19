/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */
//  8 ms 98.03 %
func permuteUnique(nums []int) [][]int {
	used := make([]bool, len(nums))
	rtn := make([][]int, 0)
	sort.Ints(nums)
	dfs(nums, 0, used, []int{}, &rtn)
	return rtn
}

func dfs(candidates []int, l int, used []bool, curr []int, ans *[][]int) {
	if l == len(candidates) {
		tmp := make([]int, l)
		copy(tmp, curr)
		(*ans) = append((*ans), tmp)
		return 
	}

	for i:=0; i<len(candidates); i++ {
		if used[i] == true {
			continue
		}
		if i > 0 && used[i-1] == true && candidates[i] == candidates[i-1] { // 去重
			// 一定注意此处的去重条件，
			// i>0保证i-1不越界，不是第一个
			// used保证前一个必须是访问过的
			// candidates保证同一层访问过的元素不再访问
			continue
		}
		used[i] = true
		curr = append(curr, candidates[i])
		dfs(candidates, l+1, used, curr, ans)
		curr = curr[:len(curr)-1]
		used[i] = false
	}
}

