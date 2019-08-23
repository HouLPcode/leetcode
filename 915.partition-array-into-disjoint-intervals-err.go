/*
 * @lc app=leetcode id=915 lang=golang
 *
 * [915] Partition Array into Disjoint Intervals
 */
// 注意，思路错误
// [1,1]
// [90,47,69,10,43,92,31,73,61,97]
func partitionDisjoint(A []int) int {
	// 遍历一遍，找到最小值位置，并且记录下来最小值前一直的最大值
	min, minindex, premax := 1000001, -1, -1
	max := -1
	for i := 0; i < len(A); i++ {
		if A[i] > max { // 注意，此处统计当前的最大值
			max = A[i]
		}
		if A[i] < min {
			min = A[i]
			minindex = i
			if max > premax { // 注意，此处更新最小值之前的最大值，即此时当前值是最小值
				premax = max
			}
		}
	}
	for premax > A[minindex] { // 此处是贪心想法，不一定正确，在minindex后面可能还有小于premax的值
		minindex++
	}
	if minindex == 0 {
		return 1
	}
	return minindex // 输出的是长度，此处索引大1正好抵消+1操作
}
