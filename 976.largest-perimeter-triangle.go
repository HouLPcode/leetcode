import "sort"

/*
 * @lc app=leetcode id=976 lang=golang
 *
 * [976] Largest Perimeter Triangle
 */
//  44 ms, faster than 73.03%
func largestPerimeter(A []int) int {
	// 构成三角形的边满足 a<=b<=c 且 a+b>c
	// 先排序，然后从右往左遍历
	// 如果 A[n-2]，A[n-1], A[n]不能构成三角形，即 A[n-2]+A[n-1] <= A[n],
	// 则A[n-3...0]不可能满足三角形， 舍去A[n]
	sort.Ints(A)
	for i := len(A) - 1; i >= 2; i-- {
		if A[i-2]+A[i-1] > A[i] {
			return A[i-2] + A[i-1] + A[i]
		}
	}
	return 0
}
