/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */
func generateParenthesis(n int) []string {
	var result []string
	help_gen(0,0,n,"",&result)
	return result
}
// 注意这里的result 需要是指针类型 *[]string
func help_gen(lused,rused,n int,tmp string, result *[]string){
	//结束条件
	if lused == n && rused == n{
		*result = append(*result,tmp)
	}
	// 思考下为什么不是左括号用完后才使用右括号
	if lused < n{
		//还有(可用
		help_gen(lused+1, rused, n, tmp + "(" ,result)
	}
	if rused < n && rused < lused{
		help_gen(lused, rused + 1, n, tmp + ")" ,result)
	}
}

