/*
 * @lc app=leetcode id=187 lang=golang
 *
 * [187] Repeated DNA Sequences
 */

 // "ABCGGCBGGGABCGGCBGGG"
func findRepeatedDnaSequences(s string) []string {
	strMap := make(map[int]int, len(s)-9)
	result := []string{}
	for i:=0; i<len(s)-9; i++ {
		strMap[str2int(s[i:i+10])]++
	}
	for k,v := range strMap{
		if v > 1{
			result = append(result, int2str(k))
		}
	}
	return result
}

func str2int(s string) int { // 不写成函数，用双指针更快速？
	// A:00 C:01 G:10 T:11
	result := 0
	for i:=0; i<10; i++{
		switch s[i]{
		case 'A': 
			result = result << 2 | 0
		case 'C':
			result = result << 2 | 1
		case 'G':
			result = result << 2 | 2
		case 'T':
			result = result << 2 | 3
		}
	}
	return result
}

func int2str(n int) string { // 不写成函数，用双指针更快速？
	// A:00 C:01 G:10 T:11
	result := ""
	for i:=0; i<10; i++{
		switch n&3{
		case 0:
			result = result + "A"
		case 1:
			result = result + "C"
		case 2:
			result += "G"
		case 3:
			result += "T"
		}
		n >>= 2
	}
	return result
}

