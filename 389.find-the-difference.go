/*
 * @lc app=leetcode id=389 lang=golang
 *
 * [389] Find the Difference
 */
func findTheDifference(s string, t string) byte {
	// hash table
	// t多一个元素, 可以有重复元素
	alphMap := make(map[byte]int, 0)
	for i:=0; i<len(s); i++ {
		alphMap[s[i]]++
	}
	var rtn byte
	for i:=0; i<len(t); i++ {
		if v,ok := alphMap[t[i]]; ok&&v!=0{
			alphMap[t[i]]--
		}else{
			rtn = t[i]
			break
		}
	}
	return rtn
}

