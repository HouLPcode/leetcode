import "sort"

/*
 * @lc app=leetcode id=767 lang=golang
 *
 * [767] Reorganize String
 */
//  (0 ms) √ Your runtime beats 100 %
// 鸽巢问题
// 普通排序，按照频次从大到小排序，隔一个放置一个
func reorganizeString(S string) string {
	// 统计S中各个字符出现的次数, 只有小写字母
	buf := make([]int, 26) // buf[字母][频次]
	for i := 0; i < len(S); i++ {
		buf[S[i]-'a']++
	}

	nums := make([]chars, 0, 26)
	for i := byte(0); i < 26; i++ {
		if buf[i] > 0 {
			nums = append(nums, chars{buf[i], 'a' + i})
		}
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i].cnt > nums[j].cnt
	})

	if nums[0].cnt > (len(S)+1)/2 { // 除了这种情况都能满足条件？？？？？？？？？？？？？
		return "" // 不可能
	}

	rtn := make([]byte, len(S))
	index := 0
	for i := 0; i < len(nums) && nums[i].cnt > 0; i++ {
		for j := 0; j < nums[i].cnt; j++ {
			if index >= len(S) {
				index = 1
			}
			//  false的情况处理
			rtn[index] = nums[i].alh
			index += 2
		}
	}
	for i := 0; i < len(S)-1; i++ {
		if rtn[i] == rtn[i+1] {
			return ""
		}
	}
	return string(rtn)
}

type chars struct {
	cnt int
	alh byte
}
