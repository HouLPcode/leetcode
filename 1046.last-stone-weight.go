import "sort"

/*
 * @lc app=leetcode id=1046 lang=golang
 *
 * [1046] Last Stone Weight
 */
// 0 ms
func lastStoneWeight(stones []int) int {
	// 循环处理，每次排序找到要处理的石头
	sort.Ints(stones)
	for len(stones) > 1 {
		l := len(stones)
		if stones[l-1] == stones[l-2] {
			stones = stones[:l-2] // 去掉两个石头
		} else {
			stones[l-2] = stones[l-1] - stones[l-2]
			stones = stones[:l-1]
			sort.Ints(stones)
		}
	}
	if len(stones) == 0 {
		return 0
	}
	return stones[0]
}
