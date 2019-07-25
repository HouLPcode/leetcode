/*
 * @lc app=leetcode id=400 lang=golang
 *
 * [400] Nth Digit
 */
// 0 ms
func findNthDigit(n int) int {
	if n < 1 {
		return -1
	}

	// 1. calculate how many digits the number has.
	digs := 1 // 最少1位数
	for ; n-countOfIntegers(digs) > 0; digs++ {
		n -= countOfIntegers(digs)
	}
	// 一共 digs位数，n为从digs位开始的数剩余的个数
	// 2. calculate what the number is.
	start := beginNumber(digs) // 从这个数开始数，比如100，1000
	if n%digs == 0 {
		start = start + n/digs - 1 // 能除尽，是上一个数
	} else {
		// 不能除尽
		start = start + n/digs
	}
	num := start // 定位到具体的数字
	// 3. find out which digit in the number is we wanted.
	if n%digs == 0 {
		return start % 10
	}
	for i := n % digs; i < digs; i++ {
		num /= 10
	}
	return num % 10
}

func countOfIntegers(digits int) int {
	s := 9
	for i := 1; i < digits; i++ {
		s *= 10
	}
	return s * digits
}

func beginNumber(digits int) int { // m位数的第一个数字， 3->100
	rtn := 1
	for i := 1; i < digits; i++ {
		rtn *= 10
	}
	return rtn
}
