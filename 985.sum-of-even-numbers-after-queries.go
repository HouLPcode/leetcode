/*
 * @lc app=leetcode id=985 lang=golang
 *
 * [985] Sum of Even Numbers After Queries
 */
// 364 ms, faster than 75.35%
func sumEvenAfterQueries(A []int, queries [][]int) []int {
	evenSum := 0
	for i := 0; i < len(A); i++ {
		if A[i]&1 == 0 {
			evenSum += A[i]
		}
	}

	ans := make([]int, len(A))
	val, index := 0, 0
	for i := 0; i < len(A); i++ {
		val = queries[i][0]
		index = queries[i][1]
		if A[index]&1 == 0 {
			// 原先是偶数
			if (A[index]+val)&1 == 0 {
				// 修改后也是偶数
				evenSum += val
			} else {
				// 修改后是奇数了
				evenSum -= A[index]
			}
		} else {
			// 原先是奇数
			if (A[index]+val)&1 == 0 {
				// 修改后是偶数
				evenSum += (val + A[index])
			} else {
				// 修改后是奇数
			}
		}
		A[index] += val // 更新A中的值
		ans[i] = evenSum
	}
	return ans
}