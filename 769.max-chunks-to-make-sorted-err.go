/*
 * @lc app=leetcode id=769 lang=golang
 *
 * [769] Max Chunks To Make Sorted
 */
// 错误思路： 从左往右遍历， 遇到比自己大的值就分块。如果最小的值在最右边，则只能分1块
// [1,2,0,3]
func maxChunksToSorted(arr []int) int {
	// -> 遍历， 遇到比自己大的值就分块
	curMax := arr[0]
	cnt := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > curMax {
			curMax = arr[i]
			cnt++
		}
	}
	cnt++
	return cnt
}
