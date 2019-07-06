/*
 * @lc app=leetcode id=942 lang=golang
 *
 * [942] DI String Match
 */
// 152 ms 78.95 % 跟最快的代码类似，不用优化
//  贪心思想，如果是I，填写当前最小值，更新最小值
//           如果是D， 填写当前最大值，更新最大值
func diStringMatch(S string) []int {
	rtn := make([]int, 0, len(S)+1)
	min, max := 0, len(S)
	for _, v := range S {
		if v == 'I' {
			rtn = append(rtn, min)
			min++
		} else {
			rtn = append(rtn, max)
			max--
		}
	}
	rtn = append(rtn, max) // 一定不要忘记最后一位，
	return rtn
}
