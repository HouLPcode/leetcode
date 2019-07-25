import (
	"sort"
	"strconv"
)

/*
 * @lc app=leetcode id=179 lang=golang
 *
 * [179] Largest Number
 */
//  0 ms
// [3,30,34]
// [0,0]
// [0]
// 拼接两个数后再比较 直接左右两边拼接成字符串后，取较大的那一个
// 通过strconv.FormatInt代替fmt.Sprint优化
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		// str1 := fmt.Sprint(nums[i])
		// str2 := fmt.Sprint(nums[j])

		// str1 := fmt.Sprint(nums[i]) + fmt.Sprint(nums[j]) // 写道一个sprint中间会出现空格
		// str2 := fmt.Sprint(nums[j]) + fmt.Sprint(nums[i])
		// if strings.Compare(str1, str2) == -1 {
		// 	return false
		// }
		// return true
		str1, str2 := strconv.FormatInt(int64(nums[i]), 10), strconv.FormatInt(int64(nums[j]), 10)
		return str1+str2 > str2+str1
	})
	rtn := ""
	for _, v := range nums {
		rtn += strconv.FormatInt(int64(v), 10)
	}
	if len(rtn) > 0 && rtn[0] == '0' {
		return "0"
	}
	return rtn
}
