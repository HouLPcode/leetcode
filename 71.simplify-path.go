/*
 * @lc app=leetcode id=71 lang=golang
 *
 * [71] Simplify Path
 */
//  0 ms
// stack 实现
 func simplifyPath(path string) string {
	params := strings.Split(path, "/")
	stack := make([]string, 0)
	for _,v := range params {
		if v == "." || v == "" {

		}else if v == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		}else {
			stack = append(stack, v)
		}
	}

	if len(stack) == 0 { // 注意，最少是 "/"
		return "/"
	}
	rtn := ""
	for _,v := range stack {
		rtn = rtn + "/" + v
	}
	return rtn
}
