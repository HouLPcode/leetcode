import "sort"

/*
 * @lc app=leetcode id=945 lang=golang
 *
 * [945] Minimum Increment to Make Array Unique
 */
// (72 ms) √ Your runtime beats 73.68 %
func minIncrementForUnique(A []int) int {
	sort.Ints(A) // 先排序
	move := 0
	for i := 0; i < len(A)-1; i++ {
		if A[i] >= A[i+1] { // 要求递增，严格 < ,所以不符合条件的需要+x
			move += (A[i] - A[i+1] + 1)
			A[i+1] = A[i] + 1
		}
	}
	return move
}
