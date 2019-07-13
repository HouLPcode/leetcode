import "strconv"

/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 */
//  0 ms
//  "0000"
// "010010"
func restoreIpAddresses(s string) []string {
	// 在s的字符的间隙中插入3个点，即[1, len(s)-2]中选出3个数，组合，一次判断
	if len(s) < 4 {
		return []string{}
	}
	// candidates = [1, len(s)-1]
	rtn := []string{}
	dfs(1, s, []int{}, &rtn)
	return rtn
}

// candidate 在那个间隙插入点，
func dfs(candidate int, s string, path []int, ans *[]string) {
	if len(path) == 3 {
		if isIP(s, path) == true { // 满足终止条件,已经放置了3个点
			ip := s[:path[0]] + "." + s[path[0]:path[1]] + "." + s[path[1]:path[2]] + "." + s[path[2]:]
			(*ans) = append((*ans), ip)
		}
		return
	}
	// 1-10
	// for i := candidate; i < len(s); i++ { // candidate
	for i := candidate; i < candidate+3 && i < len(s); i++ { // 优化，每次最多间隔3个
		// if i > candidate && s[candidate] == '0' { // not vaild 一定注意，不能出现00的情况,但是此处不知道怎么限制数组没有前缀0
		// 	continue
		// }
		// used
		path = append(path, i)
		// backtrace
		dfs(i+1, s, path, ans)
		// recover
		path = path[:len(path)-1]
	}
}

// path需要保证不重复，且范围必须在s的中间，由组合保证该条件
func isIP(s string, path []int) bool {
	str1 := s[:path[0]]
	str2 := s[path[0]:path[1]]
	str3 := s[path[1]:path[2]]
	str4 := s[path[2]:]
	// 一定注意，不能出现00的情况
	if str1[0] == '0' && len(str1) > 1 {
		return false
	}
	if str2[0] == '0' && len(str2) > 1 {
		return false
	}
	if str3[0] == '0' && len(str3) > 1 {
		return false
	}
	if str4[0] == '0' && len(str4) > 1 {
		return false
	}
	// if len(str) > 3 {} 通过组合保证每个字符串长度均在[1,3]之间
	// string not empty
	num1, _ := strconv.Atoi(str1)
	num2, _ := strconv.Atoi(str2)
	num3, _ := strconv.Atoi(str3)
	num4, _ := strconv.Atoi(str4)
	if num1 > 255 || num2 > 255 || num3 > 255 || num4 > 255 {
		return false
	}
	return true
}
