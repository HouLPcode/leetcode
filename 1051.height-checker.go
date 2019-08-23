/*
 * @lc app=leetcode id=1051 lang=golang
 *
 * [1051] Height Checker
 */
// (0 ms)
// 排序后是正确的排队位置，原顺序与排序后顺序比较，不同的位置就是站错地方了
func heightChecker(heights []int) int {
	// 桶排序
	h := make([]int, 101)
	for _, v := range heights {
		h[v]++
	}
	i := 1 // 起始身高
	cnt := 0
	for _, v := range heights {
		for h[i] == 0 { // 找到最小高度i
			i++
		}
		if i != v { // v高度的位置应该是i
			cnt++
		}
		h[i]--
	}
	return cnt
}
