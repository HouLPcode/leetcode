/*
 * @lc app=leetcode id=728 lang=golang
 *
 * [728] Self Dividing Numbers
 */
 // 暴搜
func selfDividingNumbers(left int, right int) []int {
	rtn := []int{}
	for i:=left; i<=right; i++ {
		if isSelfDividing(i) {
			rtn = append(rtn, i)
		} 
	}
	return rtn
}

func isSelfDividing(num int) bool{
	for a:=num; a!=0; {
		if (a%10)!=0 && num%(a%10) == 0{ //注意此处a%10有可能为0
			a /= 10
		}else{
			return false
		}
	}
	return true
}

