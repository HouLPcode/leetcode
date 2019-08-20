/*
 * @lc app=leetcode id=468 lang=golang
 *
 * [468] Validate IP Address
 */
// 0ms
// "1.0.1." 末尾不对
// "12..33.4"
// "2001:0db8:85a3:0:0:8A2E:0370:7334:" 段数不对
// "20EE:FGb8:85a3:0:0:8A2E:0370:7334" 不合法字符
// "2001:0db8:85a3:0:0:8A2E:0370:7334" 合法的
func validIPAddress(IP string) string {
	// ipv4
	b := isIPv4(IP)
	if b == 0 {
		return "IPv4"
	} else if b == 1 {
		return "Neither"
	}
	// ipv6
	return isIPv6(IP)
}

// 0 是 1不是 2可能是ipv6
// 只能是4段，3个点，不能忽略
func isIPv4(s string) int {
	sum := 0
	dot := 0
	if len(s) == 0 || s[0] == '.' || s[len(s)-1] == '.' {
		return 1
	}
	for i := 0; i < len(s); i++ {
		if s[i] != '.' && (s[i] < '0' || s[i] > '9') {
			return 2
		} else if s[i] == '.' {
			if sum > 255 || s[i+1] == '.' {
				return 1
			}
			dot++
			sum = 0 // 初始化
		} else {
			if sum == 0 && s[i] == '0' && i != len(s)-1 && s[i+1] != '.' { // 注意不能是末尾0
				// 前缀0
				return 1
			}
			sum *= 10
			sum += int(s[i] - '0')
		}
	}
	if sum < 0 || sum > 255 || dot != 3 { // 最后一个没有点
		return 1
	}

	return 0
}

// 只能是8段，不要忽略段数的判断
func isIPv6(s string) string {
	cnt := 0 // 统计每个字段的位数
	dot := 0
	if len(s) == 0 || s[0] == ':' || s[len(s)-1] == ':' {
		return "Neither"
	}
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'f') || (s[i] >= 'A' && s[i] <= 'F') || (s[i] >= '0' && s[i] <= '9') { // 不能忽略0-9 a-f之外的字符
			cnt++
		} else if s[i] == ':' {
			if cnt > 0 && cnt < 5 {
				cnt = 0
				dot++
			} else {
				return "Neither"
			}
		} else {
			return "Neither"
		}
	}
	if cnt > 0 && cnt < 5 && dot == 7 { // 不能忽略最后一段的判断
		return "IPv6"
	}
	return "Neither"
}
