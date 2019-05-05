/*
 * @lc app=leetcode id=461 lang=golang
 *
 * [461] Hamming Distance
 */
func hammingDistance(x int, y int) int {
	// 1. 异或运算，不同位为1
	// 2. 统计步骤1中位为1的个数
	tmp := x ^ y
	cnt := 0
	for tmp != 0{
		tmp &= tmp-1 //消除最后的1
		cnt++
	} 
	return cnt
}

