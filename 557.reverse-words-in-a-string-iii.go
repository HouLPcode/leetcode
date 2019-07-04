import "strings"

/*
 * @lc app=leetcode id=557 lang=golang
 *
 * [557] Reverse Words in a String III
 */
//  16 ms 34.36 %
func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i := 0; i < len(words); i++ {
		words[i] = recover(words[i])
	}
	rtn := ""
	for _, s := range words {
		rtn += " " + s
	}
	return rtn[1:]
}

func recover(s string) string {
	buf := []byte(s)
	for i := 0; i < len(s)/2; i++ {
		buf[i], buf[len(s)-1-i] = buf[len(s)-1-i], buf[i]
	}
	return string(buf)
}

