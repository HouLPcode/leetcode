/*
 * @lc app=leetcode id=507 lang=golang
 *
 * [507] Perfect Number
 */
//  99999991 超时
//   ✘ testcase: '1'
//   ✘ testcase: '0'
// 技巧1：遍历到num的开方即可
func checkPerfectNumber(num int) bool {
	if num <= 1 { // 注意， 0 1都是false
		return false
	}
	// 怎么找到一个数的所有因子 ？？？
	sum := 0
	for i:=1; i*i<=num; i++ { //遍历到num的开方即止
		if num%i == 0 {
			sum += i
			if i != num/i { // 避免重复计算
				sum += num/i
			}
		}
	}
	return sum == num*2
}

