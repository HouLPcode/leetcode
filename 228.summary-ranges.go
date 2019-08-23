import "strconv"

/*
 * @lc app=leetcode id=228 lang=golang
 *
 * [228] Summary Ranges
 */
// (0 ms)
// 区间遍历
func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	} else if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}
	ranges := []string{}
	s, e := 0, 1
	for e < len(nums) {
		if nums[e-1]+1 != nums[e] {
			// add range
			if s+1 == e {
				ranges = append(ranges, strconv.Itoa(nums[s]))
			} else {
				ranges = append(ranges, strconv.Itoa(nums[s])+"->"+strconv.Itoa(nums[e-1]))
			}
			s = e
		}
		e++
	}
	// add range
	if s+1 == e {
		ranges = append(ranges, strconv.Itoa(nums[s]))
	} else {
		ranges = append(ranges, strconv.Itoa(nums[s])+"->"+strconv.Itoa(nums[e-1]))
	}
	return ranges
}
