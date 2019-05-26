/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */
 // 两个指针组成子序列，map中保存出现的过的字母
func lengthOfLongestSubstring(s string) int {
	num, numMax := 0,0
	alpMap := make(map[byte]bool, 0)
	for p1,p2:=0,0; ; p1++{
		for ; p2<len(s) ; p2++{
			if _,ok := alpMap[s[p2]]; !ok{ // 字符不存在
				alpMap[s[p2]] = true
				num++
				if num > numMax{
					numMax = num
				}
			}else{
				for ; s[p1]!=s[p2]; p1++{
					delete(alpMap, s[p1])// delete不是remove
					num--
				}
				p2++ // 退出循环前右移，指向重复字母的下一个
				break //退出p2的循环
			}
		}
		if p2 == len(s){//p2已经遍历完，注意此处不是len(s)-1
			break
		}
	}
	return numMax
}
