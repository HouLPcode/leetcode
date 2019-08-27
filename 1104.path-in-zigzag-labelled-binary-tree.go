import "math"

/*
 * @lc app=leetcode id=1104 lang=golang
 *
 * [1104] Path In Zigzag Labelled Binary Tree
 */
// (0 ms)
func pathInZigZagTree(label int) []int {
	depth := int(math.Log2(float64(label))) + 1
	rtn := make([]int, depth+1)

	// 一次 label/2， 一次上一行首尾和-label/2，循环执行
	for i := depth; i > 0; i-- {
		rtn[i] = label // 不分奇偶行
		label /= 2
		i--
		rtn[i] = int(math.Pow(2, float64(i-1))+math.Pow(2, float64(i))) - 1 - label
		label /= 2
	}
	return rtn[1:]
}
