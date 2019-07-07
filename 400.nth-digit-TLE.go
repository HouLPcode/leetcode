/*
 * @lc app=leetcode id=400 lang=golang
 *
 * [400] Nth Digit
 */
// 1111111111 超时
func findNthDigit(n int) int {
	// 从1 ... n数数，同时记录数的 位数，累计到n后停止
	num := 1 //
	cnt := 1 // 注意此处是1不是0
	for cnt < n {
		num++ // 注意是num先加 ，然后统计个数，避免最终结束是num的前一个数
		cnt += calc(num)
	}

	// 结果是num中的某个数字
	for cnt != n {
		cnt--
		num /= 10
	}
	return num % 10
}

// 统计a的位数
func calc(a int) int {
	rtn := 0
	for a != 0 {
		rtn++
		a /= 10
	}
	return rtn
}
