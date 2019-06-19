/*
 * @lc app=leetcode id=216 lang=golang
 *
 * [216] Combination Sum III
 */
//  0 ms
// 位操作
func combinationSum3(k int, n int) [][]int {
	// 从1-9中挑k个数，sum为n
	rtn := [][]int{}

	for i:=1; i< (1<<9); i++ { 		// i有k个1且对应和为n
		curr := make([]int, 0, 9)
		bits, sum := 0, 0
		for j:=1; j<=9; j++ {
			if i & (1<<(uint(j-1))) != 0 {
				bits++
				sum += j
				curr = append(curr, j)
			}
		}
		if bits == k && sum == n {
			rtn = append(rtn, curr)
		}
	}
	
	return rtn
}
