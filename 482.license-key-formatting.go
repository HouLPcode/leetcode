/*
 * @lc app=leetcode id=482 lang=golang
 *
 * [482] License Key Formatting
 */
// 0ms
// memory: 4.4MB 96.67%
//  "2-4A0r7-4k"\n4
// "--a-a-a-a--"\n2
// "aaaa"\n2
// 题目理解有误，除第一组外，其他每组均有K个元素
// string的拼接通过[]byte->string加速
func licenseKeyFormatting(S string, K int) string {
	//从右往左处理
	out := make([]byte, len(S)*2) // out长度不一定小于S
	times := 0
	index := len(S)*2-1
	for i:=len(S)-1; i>=0; i--{
		if S[i] != '-' {
			if times == K{ // 放在S[i]赋值前，避免最终结果头多一个 "-"
				times = 0
				out[index] = '-'
				index--
			}
			times++
			if S[i] >= 'a'{ // convert uppercase
				out[index] = S[i]-'a'+'A'
			} else {
				out[index] = S[i]
			}
			index--
		}
	}
	return string(out[index+1:]) 	// 去除头部 -
}
