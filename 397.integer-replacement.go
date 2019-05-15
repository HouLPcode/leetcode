/*
 * @lc app=leetcode id=397 lang=golang
 *
 * [397] Integer Replacement
 */
func integerReplacement(n int) int {
	// n=0不可能变成1
	// 注意 奇数++或--的结果不一样 1234，65535
	// 什么时候++ 什么时候--
	cnt1 := 0
	for num1 := n; num1 != 1;{
		if num1 & 1 == 1{
			num1++ //
		}else{
			num1 >>= 1
		}
		cnt1++
	}

	cnt2 := 0
	for num2 :=n; num2 != 1;{
		if num2 & 1 == 1{
			num2-- //
		}else{
			num2 >>= 1
		}
		cnt2++
	}
	if cnt1 > cnt2{
		return cnt2
	}
	return cnt1
}

