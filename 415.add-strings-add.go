/*
 * @lc app=leetcode id=415 lang=golang
 *
 * [415] Add Strings
 */
//  0 ms
func addStrings(num1 string, num2 string) string {
	rtn := make([]byte, len(num1)+len(num2))//浪费了点空间
	f := byte('0') // 进位标志
	i, j := len(num1)-1, len(num2)-1
	index := len(rtn)-1
	for ; i>=0&&j>=0; i,j = i-1,j-1 {
		f,rtn[index] = add(num1[i],num2[j],f)
		index--
	}
	for ; i>=0; i--{
		f,rtn[index] = add(num1[i],'0',f)
		index--
	}
	for ; j>=0; j--{
		f,rtn[index] = add('0',num2[j],f)
		index--
	}
	if f != '0' {
		rtn[index] = f
		index-- //注意此处不能少，需要保持不管有没有首位进位，返回值都是rtn[index+1:]
	}
	return string(rtn[index+1:])
}

func add(a,b,c byte)(byte,byte) {
	return (a+b+c-'0'*3)/10+'0',(a+b+c-'0'*3)%10+'0'
}

