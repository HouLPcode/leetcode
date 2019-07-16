import "sort"

/*
 * @lc app=leetcode id=881 lang=golang
 *
 * [881] Boats to Save People
 */
//   96 ms, faster than 93.02%
func numRescueBoats(people []int, limit int) int {
	// 排序，右往左遍历，大于limit的直接计数，之后左右指针计算和后决定怎么计数
	sort.Ints(people)
	cnt := 0
	r := len(people) - 1
	l := 0
	for l < r {
		if people[l]+people[r] > limit {
			cnt++
			r--
		} else if people[l]+people[r] <= limit {
			cnt++
			l++
			r--
		}
	}
	// 退出循环时如果剩一个人，则需要一个船, l>r则表示所有人已经分配完毕
	if l == r {
		cnt++
	}
	return cnt
}
