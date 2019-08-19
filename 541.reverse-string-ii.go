/*
 * @lc app=leetcode id=541 lang=golang
 *
 * [541] Reverse String II
 */
// 0 ms
func reverseStr(s string, k int) string {
	tmp := []byte(s)
	i := 0
	for ; i+2*k-1 < len(tmp); i += 2 * k {
		reverse(tmp[i : i+k])
	}

	if i+k < len(tmp) { // 剩余的够k个
		reverse(tmp[i : i+k])
	} else { // 不够k个
		reverse(tmp[i:])
	}
	return string(tmp)
}

func reverse(buf []byte) {
	l := len(buf) // k
	for i := 0; i < l/2; i++ {
		buf[i], buf[l-1-i] = buf[l-1-i], buf[i]
	}
}
