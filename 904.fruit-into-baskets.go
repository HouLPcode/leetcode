/*
 * @lc app=leetcode id=904 lang=golang
 *
 * [904] Fruit Into Baskets
 */
//   688 ms, faster than 12.5 %
func totalFruit(tree []int) int {
	// 逐个遍历，两指针
	max := 0
	if len(tree) == 0 {
		return 0
	}
	p0 := 0
	baskets := []int{-1, -1}
	l := len(tree)
	for p0 < len(tree) {
		baskets[0] = tree[p0]
		p1 := p0 + 1
		for p1 < l {
			if tree[p1] == baskets[0] {
				p1++
			} else if baskets[1] == -1 {
				baskets[1] = tree[p1]
				p1++
			} else if tree[p1] == baskets[1] {
				p1++
			} else {
				break
			}
		}
		if p1-p0 > max {
			max = p1 - p0
		}
		//  注意，不管新的长度是否比max大，只要是计算结束，buakets都需要重置为-1
		baskets[0] = -1
		baskets[1] = -1
		for p0++; p0 < l && tree[p0] == tree[p0-1]; p0++ {
		}
	}
	return max
}
