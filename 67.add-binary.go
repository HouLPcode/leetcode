/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */

 func addBinary(a string, b string) string {
	rtn := ""
	f := byte('0') // 进位标志
	var s byte
	// 从右往左遍历
	i,j := len(a)-1, len(b)-1
	for ; i>=0&&j>=0; i,j=i-1,j-1 {
		s,f = add(a[i],b[j],f)
		rtn = string(s) + rtn
	}
	// 注意i，j 不可以是for局部变量
	for i >= 0{ // a有剩余
		s,f = add(a[i],'0',f)
		rtn = string(s) + rtn //采用[]byte能减少空间？？？
		i--
	}
	for j >= 0{ // b有剩余
		s,f = add(b[j],'0',f)
		rtn = string(s) + rtn
		j--
	}
	if f == '1' {// f有剩余
		rtn = "1" + rtn
	}
	return rtn
}

func add(a,b,f byte) (byte,byte) {
	return (a+b+f-'0'*3)%2+'0', (a+b+f-'0'*3)/2+'0'
}
