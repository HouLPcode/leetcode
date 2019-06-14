/*
 * @lc app=leetcode id=482 lang=golang
 *
 * [482] License Key Formatting
 */
//  "2-4A0r7-4k"\n4
// 题目理解有误，除第一组外，其他每组均有K个元素
func licenseKeyFormatting(S string, K int) string {
	//处理第一组
	out := ""
	for i:=0; i<len(S)&&S[i]!='-'&&i<K; i++{
		out += string(S[i])
	}
	out += "-"
	time := 0 // 一组中字母的个数
	for i:=len(out)-1; i<len(S); i++ {
		if S[i] == '-'{

		}else {
			time++
			// convert uppercase
			if S[i] >= 'a' {
				out += string(S[i]-'a'+'A')
			}else {
				out += string(S[i])
			}
		}
		if time == K {
			time = 0
			out += "-"
		}
	}
	// 去除末尾的 "-"
	if out[len(out)-1] == '-'{
		return out[:len(out)-1]
	}
	return out
}

