/*
 * @lc app=leetcode id=790 lang=golang
 *
 * [790] Domino and Tromino Tiling
 */
func numTilings(N int) int {
	if N == 1 {
		return 1
	}
	if N == 2 {
		return 2
	}
	if N == 3 {
		return 5
	}
	return numTilings(N-1) + numTilings(N-2) + 2 + numTilings(N-3) + 3
}
