/*
 * @lc app=leetcode id=645 lang=golang
 *
 * [645] Set Mismatch
 */
func findErrorNums(nums []int) []int {
	//该题有六七种解法，此处选择map处理 时间O(n),空间O(n)
	rnt := []int{0,0}
	buf := make(map[int]bool) //采用bool省空间
	for _,v := range nums{
		if _,ok := buf[v]; ok{
			rnt[0] = v
			// break 此处不能停止，否则找不到缺失数字
		}else{
			buf[v] = true
		}
	}
	for i:=1; i<=len(nums);i++{
		if _,ok := buf[i];!ok{
			rnt[1] = i
			break
		}
	}
	return rnt
}

