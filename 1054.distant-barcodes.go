/*
 * @lc app=leetcode id=1054 lang=golang
 *
 * [1054] Distant Barcodes
 */
//  (668 ms) √ Your runtime beats 71.74 %
// 相邻元素不同的排序思路：统计每个元素的频次，先排序出现次数最多的元素，从0开始排序，剩余的元素一次间隔排序
func rearrangeBarcodes(barcodes []int) []int {
	// 统计每个数字出现的次数
	buf := make(map[int]int, 0) // 数字->频次
	maxTime := 0
	maxInt := 0
	for i := 0; i < len(barcodes); i++ {
		buf[barcodes[i]]++
		if maxTime < buf[barcodes[i]] {
			maxTime = buf[barcodes[i]]
			maxInt = barcodes[i]
		}
	}

	rtn := make([]int, len(barcodes))
	// 先排序频次最多的数字，从0开始排
	// 肯定有解，则最大频次不会超过一半， 奇数5个最多3个
	// 填充频次最多的数字
	index := 0 // 从0开始填充
	for i := 0; i < maxTime; i++ {
		rtn[index] = maxInt
		index += 2
	}

	// 填充剩余的数字
	delete(buf, maxInt)
	for k, v := range buf {
		for i := 0; i < v; i++ {
			if index >= len(barcodes) {
				index = 1 //从头开始填充奇数位
			}
			rtn[index] = k
			index += 2
		}
	}
	return rtn
}
