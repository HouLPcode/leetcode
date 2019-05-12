/*
 * @lc app=leetcode id=179 lang=golang
 *
 * [179] Largest Number
 */
func largestNumber(nums []int) string {
	// 直接左右两边拼接成字符串后，取较大的那一个
	// 典型错误，0存在的时候局部最优不是最终最优
	str := ""
	for _,v := range nums{
		if strings.Compare(str+fmt.Sprint(v),fmt.Sprint(v)+str) > 0{
			str = str+fmt.Sprint(v)
		} else{
			str = fmt.Sprint(v) + str
		}
	}
	return str
}
