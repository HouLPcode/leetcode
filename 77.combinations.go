/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */
//  152 ms 81.78 % 
//  1-n中选k个数的所有组合数
func combine(n int, k int) [][]int {
	rtn := make([][]int, 0)
	dfs(1, n, k, []int{}, &rtn)
	return rtn
}

func dfs(s, n, k int ,curr []int, ans *[][]int) {
	if k == 0 {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		(*ans) = append((*ans), tmp)
		return 
	}

	for i:=s; i<=n; i++ {
		curr = append(curr, i)
		dfs(i+1, n, k-1, curr, ans)
		curr = curr[:len(curr)-1]
	}
}

