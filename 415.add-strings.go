/*
 * @lc app=leetcode id=415 lang=golang
 *
 * [415] Add Strings
 */
//  err,麻烦
func addStrings(num1 string, num2 string) string {
	rtn := make([]byte, len(num1)+len(num2))//浪费了点空间
	copy(rtn[len(num2):], num1)
	j := len(num2)-1
	f := byte('0') // 进位标志
	index := len(rtn)-1
	for ; j>=0; j-- {
		f,rtn[index] = (num2[j]+rtn[index]+f-'0'*3)/10+'0',(num2[j]+rtn[index]+f-'0'*3)%10+'0'
		index--
	}
	for f > '0' {
		// rtn中默认值需要是'0',创建时是0，需要处理下
		f,rtn[index] = (rtn[index]+f-'0'*2)/10+'0',(rtn[index]+f-'0'*2)%10+'0'
		index--
	}
	return string(rtn[index+1:])
}

