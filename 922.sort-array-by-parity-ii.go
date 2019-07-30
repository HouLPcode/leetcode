/*
 * @lc app=leetcode id=922 lang=golang
 *
 * [922] Sort Array By Parity II
 */
//   × testcase: '[3,1,4,2]'
// 奇数位为奇数，偶数位为偶数
func sortArrayByParityII(A []int) []int {
	for i := 0; i < len(A); i++ {
		if i&1 == A[i]&1 {
			continue
		}
		j := i + 1 // 寻找最近的奇偶数,此处可以优化下
		for j < len(A) && i&1 != A[j]&1 {
			j++
		}
		A[i], A[j] = A[j], A[i]
		// i = j, 注意从i继续，不能从j往后走
	}
	return A
}
