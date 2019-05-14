/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */
func lengthOfLastWord(s string) int {
	// "a " 末尾空格要忽略
	cnt := 0
	for _,v := range []byte(s){
		if v == byte(' '){
			cnt = 0
		}else{
			cnt++
		}
	}
	return cnt
}

