/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */
func lengthOfLastWord(s string) int {
	// "a " 末尾空格要忽略
	pri,cnt := 0,0//通过pri暂存之前的结果，避免末尾多个0的清空
	for _,v := range []byte(s){
		if v == byte(' '){
			if cnt == 0{
				continue
			}else{
				pri,cnt = cnt,0
			}
		}else{
			cnt++
		}
	}
	if cnt > 0{
		return cnt
	}
	return pri
}

