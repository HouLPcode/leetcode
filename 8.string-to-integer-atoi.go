/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */
func myAtoi(str string) int {
	i:=0
	for i < len(str) && str[i] == byte(' '){ //删除前缀空格,全是空格会导致越界
		i++
	}
	if i == len(str){
		return 0 //全是空格
	}
	str = str[i:]
	mflag := false //+-
	if str[0] == byte('-'){
		mflag = true // -
		str = str[1:]
	}else if str[0] == byte('+'){
		mflag = false
		str = str[1:]
	}

	i=0
	for i < len(str) && int(str[i]-byte('0')) >= 0 && int(str[i]-byte('0')) <= 9{//数字
		i++
	}
	if i == 0{
		return 0 // 字母前缀
	}
	str = str[:i]// 只剩字母
	//int32 max min
	const MinInt int = -(1 << (32-1))
	const MaxInt int = (1 << (32-1)) - 1

	cnt := 0
	if mflag{
		// -
		for _,v := range []byte(str){
			if MaxInt  - cnt*10 < int(v-byte('0')) - 1{
				return MinInt
			}
			cnt = cnt*10 + int(v-byte('0'))
		}
		cnt = -cnt
	}else{
		for _,v := range []byte(str){
			if MaxInt - cnt*10 < int(v-byte('0')){ // 越界检查
				cnt = MaxInt
				break
			}
			cnt = cnt*10 + int(v-byte('0'))
		}
	}
	
	return cnt
}

