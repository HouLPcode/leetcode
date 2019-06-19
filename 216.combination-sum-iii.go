/*
 * @lc app=leetcode id=216 lang=golang
 *
 * [216] Combination Sum III
 */
//  0 ms
// 一定注意剪枝的时候放在dfs的for循环中
func combinationSum3(k int, n int) [][]int {
	// 从1-9中挑k个数，sum为n
	rtn := [][]int{}
	dfs(1, k, n, []int{}, &rtn)
	return rtn
}

func dfs(s, k, target int, curr []int, ans*[][]int) {
	if target == 0 && k == 0 {
		tmpdata := make([]int, len(curr))
		copy(tmpdata, curr)
		(*ans) = append((*ans), tmpdata)
		return 
	}

	for i:=s; i <= 9; i++ {
		// 注意剪枝放在for循环中
		if i > target || k == 0{
			return 
		}
		curr = append(curr, i)
		dfs(i+1, k-1, target-i, curr, ans)
		curr = curr[:len(curr)-1]
	}
}

