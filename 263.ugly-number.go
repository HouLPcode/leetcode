/*
 * @lc app=leetcode id=263 lang=golang
 *
 * [263] Ugly Number
 */
func isUgly(num int) bool {
	// 注意 0%x = 0
	if num == 0{
		return false
	}
    for num%5 == 0{
		num /= 5
	}
	for num%3 == 0{
		num /= 3
	}
	for num%2 == 0{
		num /= 2
	}
	return num == 1
}

