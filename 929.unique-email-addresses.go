/*
 * @lc app=leetcode id=929 lang=golang
 *
 * [929] Unique Email Addresses
 */
//  0 ms
//  ["testemail@leetcode.com","testemail1@leetcode.com","testemail+david@lee.tcode.com"]
//  strings + map
func numUniqueEmails(emails []string) int {
	rtn := make(map[string]bool, 0)
	for _, v := range emails {
		rtn[conv(v)] = true
	}
	return len(rtn)
}

func conv(s string) string {
	rtn := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			continue
		} else if s[i] == '+' {
			for ; s[i+1] != '@'; i++ {
			} // 注意此处是i+1,避免外层for多计算
		} else if s[i] == '@' {
			rtn = append(rtn, []byte(s[i:])...)
			break
		} else {
			rtn = append(rtn, s[i])
		}
	}
	return string(rtn)
}

