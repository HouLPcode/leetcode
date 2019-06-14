/*
 * @lc app=leetcode id=482 lang=golang
 *
 * [482] License Key Formatting
 */
//  "2-4A0r7-4k"\n4
// "--a-a-a-a--"\n2
// 题目理解有误，除第一组外，其他每组均有K个元素
func licenseKeyFormatting(S string, K int) string {
	//从右往左处理
	out := ""
	times := 0
	for i:=len(S)-1; i>=0; i--{
		if times == K{ // 放在S[i]赋值前，避免最终结果头多一个 "-"
			times = 0
			out = "-" + out
		}
		if S[i] != '-' {
			times++
			if S[i] >= 'a'{ // convert uppercase
				out = string(S[i]-'a'+'A') + out
			}else {
				out = string(S[i]) + out
			}
		}
	}
	return out 	// 去除末尾 -
}

