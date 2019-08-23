/*
 * @lc app=leetcode id=825 lang=golang
 *
 * [825] Friends Of Appropriate Ages
 */
//  (32 ms) √ Your runtime beats 96.43 %
func numFriendRequests(ages []int) int {
	// 按照 age 计数， 然后 依次比较 时间O(120*120)
	age := make([]int, 121)
	for _, v := range ages {
		age[v]++
	}
	cnt := 0
	for i := 1; i <= 120; i++ {
		if age[i] == 0 {
			continue
		} else {
			for j := 1; j <= 120; j++ {
				if age[j] == 0 {
					continue
				}
				if i == j {
					if isOK(i, i) { // 一定注意，此处是排列数，不是组合数
						cnt += age[i] * (age[i] - 1)
					}
				} else if isOK(i, j) {
					cnt += age[i] * age[j]
				}
			}
		}
	}
	return cnt
}

func isOK(age1, age2 int) bool {
	if (age2 <= age1/2+7) ||
		(age2 > age1) ||
		(age2 > 100 && age1 < 100) {
		return false
	}
	return true
}
