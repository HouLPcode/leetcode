/*
 * @lc app=leetcode id=1128 lang=golang
 *
 * [1128] Number of Equivalent Domino Pairs
 */
// (28 ms) √ Your runtime beats 93.55 %
func numEquivDominoPairs(dominoes [][]int) int {
	// 保证骨牌小数字在前面，大数字在后面
	nums := [10][10]int{} // 缓存次数
	for i := 0; i < len(dominoes); i++ {
		if dominoes[i][0] >= dominoes[i][1] {
			nums[dominoes[i][1]][dominoes[i][0]]++
		} else {
			nums[dominoes[i][0]][dominoes[i][1]]++
		}
	}
	cnt := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if nums[i][j] > 1 { // 组合数
				cnt += nums[i][j] * (nums[i][j] - 1) / 2
			}
		}
	}
	return cnt
}
