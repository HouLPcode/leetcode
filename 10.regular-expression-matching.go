/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */
// 8ms 49.44%
// 还有一种 DP方式实现
// 递归法 实现如下
// 假设模式中没有 * ，代码实现如下
// func isMatch(s string, p string) bool {
// 	if len(p) == 0 {
// 		return len(s) == 0
// 	}

// 	// 匹配第一个字符
// 	// 如果 s 的长度是0 则肯定不匹配
// 	// s和p的首字母匹配
// 	// p的首字母是通配符 .
// 	firstMatch := len(s) != 0 && (s[0] == p[0] || p[0] == '.')
// 	return firstMatch && isMatch(s[1:], p[1:])
// }

// 如果模式中有 * ，则出现在p[1]位置上， 首位是*是不合法的.
// * 的存在有二种功能, 前面的字符出现0次, 前面的字符出现一次或多次
// 1. 如果 s[0]!=p[0],则前面字符出现0次, 忽略p中的该部分,即 p = p[2:]
// 2. 如果 s[0]=p[0], 则前面的字符出现一次或多次. s可以右移一个字符. 表示 * 可以匹配任意多个 s[0]

// 改造后的代码如下
// "aaa" "a*a"
func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	if len(p) > 1 && p[1] == '*' { // 有*
		firstMatch := len(s) > 0 && (s[0] == p[0] || p[0] == '.') // 可以提取到if 外边, if...else中都有这句
		// 一定注意,此处不是二选一,应该是两种情况都需要处理, 因为即使第一个字符匹配上了,也可以是p[2]的匹配(s[0]==p[2])
		// if firstMatch {                                           // 即 s[0] == p[0]
		// 	return isMatch(s[1:], p)
		// } else {
		// 	return isMatch(s[:], p[2:])
		// }
		return isMatch(s[:], p[2:]) || (firstMatch && isMatch(s[1:], p))
	} else { // 没有 *
		firstMatch := len(s) > 0 && (s[0] == p[0] || p[0] == '.')
		return firstMatch && isMatch(s[1:], p[1:])
	}
	return false // 此处不会执行
}
