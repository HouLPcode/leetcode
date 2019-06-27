/*
 * @lc app=leetcode id=1018 lang=golang
 *
 * [1018] Binary Prefix Divisible By 5
 */
//  (320 ms) 98.26 % 
// 典型错误， 数组过长会导致生成的int越界
func prefixesDivBy5(A []int) []bool {
	rtn := make([]bool, len(A))
	num := 0
	for k,v := range A  {
		num = (num<<1 + v) % 5 // 一定注意该步骤，避免越界
		if num%5 == 0 {
			rtn[k] = true
			// num = 0 //能被5整除，舍弃该值，在num计算的之前同样会产生越界，如123123123123...   *10,前面越界导致计算结果不准确
		}
	}
	return rtn
}

