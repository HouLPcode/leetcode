/*
 * @lc app=leetcode id=43 lang=golang
 *
 * [43] Multiply Strings
 */
//  4 ms 56.2 %
//  "9133"\n"0"
// "0"\n"52"
//  按照乘法计算公式，先乘后加，得出结果
func multiply(num1 string, num2 string) string {
	// string -> byte, 从右往左按照乘法运算计算
	if len(num1) < len(num2) {
		return multiply(num2, num1) // 注意此处，避免出现“0000”
	}
	rtn := "0" //注意初始化为0，不要是“”，
	for j := len(num2) - 1; j >= 0; j-- {
		mult := numtbytes(num1, num2[j])
		rtn = addbytes(rtn, mult)
		num1 += "0" //注意，每计算一次后需要*10
	}
	return rtn
}

// a * b
func numtbytes(a string, b byte) string {
	flag := byte('0') // 进位
	rtn := make([]byte, 0)
	if b == '0' {
		return "0" // 一定注意，乘以0直接返回“0”，避免出现 "0000"的情况
	}
	for i := len(a) - 1; i >= 0; i-- {
		data := (a[i]-'0')*(b-'0') + (flag - '0')
		rtn = append(rtn, data%10+'0')
		flag = data/10 + '0'
	}
	if flag != '0' {
		rtn = append(rtn, flag)
	}
	// swap
	l := len(rtn)
	for i := 0; i < l/2; i++ {
		rtn[i], rtn[l-1-i] = rtn[l-i-1], rtn[i]
	}
	return string(rtn)
}

func addbytes(a, b string) string {
	if a == "" {
		return b
	} else if b == "" {
		return a
	}
	al, bl := len(a)-1, len(b)-1
	sum, flag := byte(0), byte(0)
	rtn := make([]byte, 0)
	for al >= 0 && bl >= 0 {
		sum = (a[al] - '0') + (b[bl] - '0') + (flag)
		rtn = append(rtn, sum%10+'0')
		flag = sum / 10
		al--
		bl--
	}
	for al >= 0 {
		sum = (a[al] - '0') + (flag)
		rtn = append(rtn, sum%10+'0')
		flag = sum / 10
		al--
	}
	for bl >= 0 {
		sum = (b[bl] - '0') + (flag)
		rtn = append(rtn, sum%10+'0')
		flag = sum / 10
		bl--
	}
	if flag == byte(1) { // 两个数加法最多进一位
		rtn = append(rtn, '1')
	}
	// swap
	l := len(rtn)
	for i := 0; i < l/2; i++ {
		rtn[i], rtn[l-1-i] = rtn[l-i-1], rtn[i]
	}
	return string(rtn)
}
