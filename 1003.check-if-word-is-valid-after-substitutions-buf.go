/*
 * @lc app=leetcode id=1003 lang=golang
 *
 * [1003] Check If Word Is Valid After Substitutions
 */
//  88 ms, faster than 5.00%
// 可以通过栈进行优化
// 通过在缓存中不断的删除“abc”，判断最后是不是空字符串
func isValid(S string) bool {
	tmp := []byte(S)
	for i := 0; i < len(tmp)-2; {
		if tmp[i] == 'a' && tmp[i+1] == 'b' && tmp[i+2] == 'c' {
			tmp = append(tmp[:i], tmp[i+3:]...)
			i = 0 // i从头开始, 一定注意for循环中有没有i++
		} else {
			i++
		}
	}
	return len(tmp) == 0
}
