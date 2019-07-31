import "sort"

/*
 * @lc app=leetcode id=274 lang=golang
 *
 * [274] H-Index
 */
// (0 ms) √ Your runtime beats 100 %
//   × testcase: '[]'
//   × testcase: '[0]'
// [9,5]
//   × testcase: '[0,1]'
func hIndex(citations []int) int {
	// 按照从大到小排序， 右往左找nums[i] >= i ， 折半查找？？？
	l := len(citations)
	if l == 0 {
		return 0
	}
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	// 折半尝试
	s, e := 0, l-1
	for s <= e {
		mid := (e-s)/2 + s
		if citations[mid] >= mid+1 {
			s = mid + 1
		} else {
			e = mid - 1
		}
	}
	return s // 文章数是索引+1
}
