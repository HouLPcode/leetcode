/*
 * @lc app=leetcode id=405 lang=golang
 *
 * [405] Convert a Number to Hexadecimal
 */
//  0 ms
//  库函数 strconv.FormatInt(123, 10)
//  注意负数的处理，补码表示
// &替代% >>替代/
func toHex(num int) string {
	if num == 0 { // 一定注意0的单独处理
		return "0"
	}
	rtn := ""
	buf := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}

	rtn = buf[num&0xf] + rtn
	num = num >> 4 & 0x0fffffff // 注意此处的处理，主要针对负数的情况，将高4位补的1置0，题目给出32bit
	for num != 0 {
		rtn = buf[num&0xf] + rtn
		num >>= 4
	}

	return rtn
}

