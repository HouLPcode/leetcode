/*
 * @lc app=leetcode id=1004 lang=golang
 *
 * [1004] Max Consecutive Ones III
 */
// 60 ms, faster than 81.82%
// [0,0,0,1] 4
// [left,right] 维护一个窗口，每次右移right
func longestOnes(A []int, K int) int {
	if len(A) == 0 {
		return 0
	}
	left, right := 0, 0 // 数组第一个值是0的处理？？？
	cnt0 := 0           // 当前窗口0的个数
	maxCnt := 0         // 窗口最大值
	for ; right < len(A); right++ {
		if A[right] != 1 {
			// 窗口中可以继续填充0
			cnt0++ // 此处是0，需要先更新，再和K值比较
			if cnt0 > K {
				if right-left > maxCnt { // 更新一次maxCnt统计最大窗口值
					maxCnt = right - left
				}
				for A[left] == 1 { // 找到left的第一个0，这样right才能进入窗口
					left++
				}
				left++
			}
		}
	}
	// 注意，退出循环后需要更新一次maxCnt
	if right-left > maxCnt { // 统计最大窗口值
		maxCnt = right - left
	}
	return maxCnt
}
