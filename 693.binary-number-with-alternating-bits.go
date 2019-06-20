/*
 * @lc app=leetcode id=693 lang=golang
 *
 * [693] Binary Number with Alternating Bits
 */
//  0 ms
// 0位开始，逐位判断
func hasAlternatingBits(n int) bool {
	arr := []int{1,0}
	if n&1 == 0 {
		for i:=0; n!=0; i++{
			if n&1 != arr[(i+1)%2] {
				return false
			}
			n >>= 1
		}
	}else {
		for i:=0; n!=0; i++{
			if n&1 != arr[i%2] {
				return false
			}
			n >>= 1
		}
	}
	return true
}

