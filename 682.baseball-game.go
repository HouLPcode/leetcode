import "strconv"

/*
 * @lc app=leetcode id=682 lang=golang
 *
 * [682] Baseball Game
 */
//  0 ms
func calPoints(ops []string) int {
	stack := make([]int, 0, len(ops))
	for _, v := range ops { // 所有字符转换成成绩数值
		// 不考虑非法操作
		if v == "+" {
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		} else if v == "D" {
			stack = append(stack, 2*stack[len(stack)-1])
		} else if v == "C" {
			stack = stack[:len(stack)-1]
		} else {
			dig, _ := strconv.Atoi(v)
			stack = append(stack, dig)
		}
	}
	sum := 0
	for _, v := range stack { // 求和
		sum += v
	}
	return sum
}

