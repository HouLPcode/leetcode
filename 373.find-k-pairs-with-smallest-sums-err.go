/*
 * @lc app=leetcode id=373 lang=golang
 *
 * [373] Find K Pairs with Smallest Sums
 */
// [1,1,2]\n[1,2,3]\n10
// [1,2]\n[3]\n3
// 注意，有可能所有组合数比k小
// 思路错误，比如顺序依次是(0,0),(1,0),(2,0),(3,0),(3,1) 下面的选项不一定是(4,1)或(3,2),有可能是(1,1),(2,1),因为这两个数还没有比较过
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	l1, l2 := len(nums1), len(nums2)
	rtn := make([][]int, 0, k)
	if l1 == 0 || l2 == 0 {
		return [][]int{}
	}

	i, j := 0, 0
	rtn = append(rtn, []int{nums1[i], nums2[j]})

	for len(rtn) < k && i+1 < l1 && j+1 < l2 {
		if nums1[i+1]+nums2[j] < nums1[i]+nums2[j+1] {
			rtn = append(rtn, []int{nums1[i+1], nums2[j]})
			i++
		} else {
			rtn = append(rtn, []int{nums1[i], nums2[j+1]})
			j++
		}
	}

	if len(rtn) == k {
		return rtn
	}

	if i+1 == l1 {
		if j == 0 { // 没有进入 for上面的for循环
			j++
		}
		for l := len(rtn); l <= k && j < l2; l++ {
			rtn = append(rtn, []int{nums1[i], nums2[j]})
			j++
		}
	}
	if j+1 == l2 {
		if i == 0 { // 没有进入 for上面的for循环
			i++
		}
		for l := len(rtn); l <= k && i < l1; l++ {
			rtn = append(rtn, []int{nums1[i], nums2[j]})
			i++
		}
	}
	return rtn
}
