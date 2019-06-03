/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] ZigZag Conversion
 */
//  8ms T:87.03
//  "A"\n1
 func convert(s string, numRows int) string {
	rtn := ""
	length := len(s)
	if numRows == 1{
		return s
	}
	//row 0
	for n:=0; n<length; {
		rtn += string(s[n])
		n += (numRows-1)*2 //numRows==1,死循环
	}

	for row:=1; row<numRows-1; row++ {
		for n:=row; n<length; {
			rtn += string(s[n])
			n += (numRows-row-1)*2
			if n >= length { // 注意此处的判断，避免下行访问的越界
				break
			}
			rtn += string(s[n])
			n += row*2
		}
	}

	//row numRows
	for n:=numRows-1; n<length; {
		rtn += string(s[n])
		n += (numRows-1)*2
	}
	return rtn
}
