/*
 * @lc app=leetcode id=925 lang=golang
 *
 * [925] Long Pressed Name
 */
//  0 ms
//  "pyplrz"\n"ppyypllr"
func isLongPressedName(name string, typed string) bool {
	// 遍历name的时候检查typed是否符合要求
	n, t := 0, 0
	for n < len(name) && t < len(typed) {
		if name[n] == typed[t] {
			n++
			t++
		} else if name[n] != typed[t] && t > 0 && typed[t-1] == typed[t] { // 按键重复
			t++
		} else {
			return false
		}
	}
	// 一定注意，退出循环不保证true，需要对后续剩余字符串进行判断
	if n == len(name) {
		for t < len(typed) && t > 0 && typed[t] == typed[t-1] {
			t++
		}
		return t == len(typed)
	}
	return false
}
