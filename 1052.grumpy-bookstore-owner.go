/*
 * @lc app=leetcode id=1052 lang=golang
 *
 * [1052] Grumpy Bookstore Owner
 */
//  (32 ms)  √ Your runtime beats 91.76 %
// 滑动窗口题目
// 第一遍，扫描所有高兴顾客的数目
// 第二遍，维护一个大小X的窗口，统计窗口内不高兴顾客的最大值
func maxSatisfied(customers []int, grumpy []int, X int) int {
	customer0 := 0 // 原始高兴
	// 滑动窗口
	left, right := 0, 0
	customer1 := 0
	maxCustomer1 := 0
	for ; right < len(customers); right++ {
		if grumpy[right] == 0 {
			customer0 += customers[right]
		} else {
			customer1 += customers[right]

			for right-left+1 > X { // 此处是for循环，不是if，此时的窗口可能已经远远大于X
				if grumpy[left] == 1 {
					customer1 -= customers[left]
				}
				left++
			}

			if customer1 > maxCustomer1 {
				maxCustomer1 = customer1
			}
		}
	}
	return maxCustomer1 + customer0
}
