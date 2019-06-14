/*
 * @lc app=leetcode id=443 lang=golang
 *
 * [443] String Compression
 */
//  ["a","a","a","b","b","a","a"]
// 注意需要保持字母的顺序
func compress(chars []byte) int {
	// key-value形式，如果v=1则只保留k
	charMap := make([]int, 126-35+1)
	for i:=0; i<len(chars); i++ {
		charMap[chars[i]-35]++
	}
	index := 0
	for k,v := range charMap {
		if v == 1 {
			chars[index] = byte(k+35)
			index++
		}else if v > 1{
			chars[index] = byte(k+35)
			index++
			//注意v大于个位数的处理
			times := fmt.Sprint(v)
			for i:=0; i<len(times); i++ {
				chars[index] = times[i]
				index++
			}
		}
	}
	return index //此处的index比实际大1，返回的是长度，所以不用-1处理
}

