/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */
func lengthOfLastWord(s string) int {
	// "a " 末尾空格要忽略
	cnt := 0
	l := len(s)-1
	//跳过末尾的n多空格
	for ;l>=0&&s[l]==byte(' ');l--{}
	for ;l>=0&&s[l]!=byte(' ');l--{
		cnt++
	}

	return cnt
}

