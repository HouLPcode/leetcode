/*
 * @lc app=leetcode id=888 lang=golang
 *
 * [888] Fair Candy Swap
 */
// (72 ms) √ Your runtime beats 63.89 %
func fairCandySwap(A []int, B []int) []int {
	// 计算 sumA, sumB,同时缓存A
	// 然后遍历 |A-B| = |sumA-sumB|/2
	mapA := make(map[int]bool)
	sumA := 0
	for _, v := range A {
		sumA += v
		mapA[v] = true
	}
	sumB := 0
	for _, v := range B {
		sumB += v
	}
	sub := (sumA - sumB) / 2
	for _, v := range B {
		if mapA[v+sub] == true {
			return []int{v + sub, v}
		}
	}
	return []int{}
}
