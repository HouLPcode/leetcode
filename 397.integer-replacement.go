/*
 * @lc app=leetcode id=397 lang=golang
 *
 * [397] Integer Replacement
 */
 // 三个测试集 1234(14)，65535(17)，10000(16),尤其是10000
func integerReplacement(n int) int {
	// n=0不可能变成1
	// 注意 奇数++或--的结果不一样 1234(14)，65535(17)，10000
	// 什么时候++ 什么时候--
	
	// 递归试试
	cnt := 0
	for  n != 1{
		if n & 1 == 1{
			// 注意此处是cnt的累加
			cnt = cnt + 1 + min(integerReplacement(n+1),integerReplacement(n-1))
			break
		}else{
			n >>= 1
			cnt++
		}
	}
	return cnt
}

func min(a,b int)int{
	if a < b{
		return a
	}
	return b
}

