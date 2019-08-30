/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 */
// 0 ms
func numTrees(n int) int {
	buf := make(map[int]int)
	buf[0] = 1
	return numsBuf(n, buf)
}

func numsBuf(n int, buf map[int]int) int {
	if n == 0 {
		return 1
	}
	if buf[n] > 0 {
		return buf[n]
	}
	sum := 0
	for l := 0; l < n; l++ {
		// 左节点的个数
		sum += numsBuf(l, buf) * numsBuf(n-l-1, buf)
	}
	buf[n] = sum
	return sum
}
