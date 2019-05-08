/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 */
func numDecodings(s string) int {
	// f(i) 包含第i的  解码方式数
	// 注意i位是 0 ，则前1位必是 1或2，f(i) = f(i-2)
	// f(i) = f(i-1) + [ f(i-2) , 条件 组合在10-26之间]
	// string 取中间字符， s[5]	
	//s 非空

	// 注意 "0",编码方式为0
	// 子串"00"的存在，编码方式为0
	// "01" -> 0
	// "230" -> 0
	// 注意 此处不仅仅是"0",还包括 "00","01","02"...
	if string(s[0]) == "0"{ //类型转换
		return 0
	}
	pri1,pri2 := 1,1//注意不是 0,1
	for i:=1; i<len(s); i++{
		// 注意，遍历的时候要从00-99逐情况分析，特别注意不能编码的情况 
		v,_ := strconv.Atoi(s[i-1:i+1])
		if v == 10 || v == 20{//f(i-2)
			pri1,pri2 = pri2,pri1
		}else if v%10 == 0 { // 00,30,40,50,60,70,80,90
			return 0
		}else if v < 10 || v > 26{ //f(i-1)
			pri1,pri2 = pri2,pri2
		}else{
			pri1,pri2 = pri2,pri1+pri2
		}
	}
	return pri2
}

