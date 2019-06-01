/*
 * @lc app=leetcode id=884 lang=golang
 *
 * [884] Uncommon Words from Two Sentences
 */
 // 技巧，A或B中只出现1次，则把所有word存储在一个map中，最后返回个数为1的word
import "strings"
 func uncommonFromSentences(A string, B string) []string {
    wordA := strings.Fields(A)
    wordB := strings.Fields(B)	
	count := make(map[string]int, 0)
	for _,v := range wordA{
		count[v]++
	}
	for _,v := range wordB{
		count[v]++
	}
	rtn := []string{}
	for k,v := range count{
		if v == 1{
			rtn = append(rtn, k)
		}
	}
	return rtn
}

