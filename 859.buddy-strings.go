/*
 * @lc app=leetcode id=859 lang=golang
 *
 * [859] Buddy Strings
 */
//  0 ms
//  一定注意：两个字符串一样和两个字母替换的场景
// "abcaa"\n"abcbb"   "aa"\n"aa"
func buddyStrings(A string, B string) bool {
	// A B 有且仅有两个字符不一样  但是"aa"交换后仍是"aa"的情况怎么处理
	lenA, lenB := len(A), len(B)
	if lenA != lenB {
		return false
	}
	// cnt := 0 // 统计不同元素的个数, 且需要两个不同元素是互换的
	diff := make([]int,0)
	for i:=0; i<lenA; i++ {
		if A[i] != B[i] {
			diff = append(diff, i)
		}
	}
	if len(diff) == 0 { // A,B字符串一样
		// A中存在某元素个数大于等于2，则成功
		alph := make(map[byte]int, 0)
		for i:=0; i<lenA; i++ {
			alph[A[i]]++
			if alph[A[i]] > 1{
				return true
			}
		}
	}
	if len(diff) == 2{
		return A[diff[0]] == B[diff[1]] && A[diff[1]] == B[diff[0]]
	}
	return false
}

