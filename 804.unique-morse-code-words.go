/*
 * @lc app=leetcode id=804 lang=golang
 *
 * [804] Unique Morse Code Words
 */
// 0 ms, faster than 100.00%
// 数组缓存莫斯码，结果存储在map中，然会map长度
func uniqueMorseRepresentations(words []string) int {
	morse := [][]byte{
		[]byte(".-"),
		[]byte("-..."),
		[]byte("-.-."),
		[]byte("-.."),
		[]byte("."),
		[]byte("..-."),
		[]byte("--."),
		[]byte("...."),
		[]byte(".."),
		[]byte(".---"),
		[]byte("-.-"),
		[]byte(".-.."),
		[]byte("--"),
		[]byte("-."),
		[]byte("---"),
		[]byte(".--."),
		[]byte("--.-"),
		[]byte(".-."),
		[]byte("..."),
		[]byte("-"),
		[]byte("..-"),
		[]byte("...-"),
		[]byte(".--"),
		[]byte("-..-"),
		[]byte("-.--"),
		[]byte("--..")}
	rtn := make(map[string]bool, 0)
	tmp := []byte{}
	for i := 0; i < len(words); i++ {
		tmp = tmp[:0]
		for j := 0; j < len(words[i]); j++ {
			tmp = append(tmp, []byte(morse[words[i][j]-'a'])...)
		}
		rtn[string(tmp)] = true
	}
	return len(rtn)
}
