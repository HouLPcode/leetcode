/*
 * @lc app=leetcode id=566 lang=golang
 *
 * [566] Reshape the Matrix
 */
// (48 ms) √ Your runtime beats 97.6 %
func matrixReshape(nums [][]int, r int, c int) [][]int {
	rows, cols := len(nums), len(nums[0])
	if rows*cols != r*c {
		return nums
	}
	ans := make([][]int, r)
	for i := 0; i < r; i++ {
		ans[i] = make([]int, c)
	}

	r1, c1 := 0, 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if c1 == c {
				r1++
				c1 = 0
			}
			ans[r1][c1] = nums[i][j]
			c1++
		}
	}
	return ans
}
