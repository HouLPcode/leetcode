/*
 * @lc app=leetcode id=443 lang=golang
 *
 * [443] String Compression
 */
//  ["a","a","a","b","b","a","a"]
// 注意需要保持字母的顺序
func compress(chars []byte) int {
	// key-value形式，如果v=1则只保留k
	index := 0
	for cur:=0; cur<len(chars);  {
		// 先存字母
		chars[index] = chars[cur]
		index++
		// 再存次数
		times := 1
		for next:=cur+1; next<len(chars)&&chars[cur]==chars[next];next++{
			times++
		}
		// 更新当前索引cur
		cur += times
		if times > 1 {
			tmp := fmt.Sprint(times)
			for i:=0; i<len(tmp); i++ {
				chars[index] = tmp[i]
				index++
			} 
		}
	}
	return index //此处的index比实际大1，返回的是长度，所以不用-1处理
}

