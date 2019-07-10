/*
 * @lc app=leetcode id=658 lang=golang
 *
 * [658] Find K Closest Elements
 */
//  188 ms 81.61 %
// 二分找到 x 应该在的位置
//-----不用这个思路 （然后左右遍历，找最小差值放入结果中）
// 然后从x位置左扩展k，右扩展k，从两个往里面夹，得到最终结果
func findClosestElements(arr []int, k int, x int) []int {
	if x < arr[0] { // 加速
		return arr[:k]
	} else if arr[len(arr)-1] < x { // 加速
		return arr[len(arr)-k:]
	}
	s, e := 0, len(arr)-1
	for s <= e {
		mid := (e-s)/2 + s
		if arr[mid] == x {
			s = mid
			break
		} else if arr[mid] > x {
			e = mid - 1
		} else {
			s = mid + 1
		}
	}
	// 不管arr[s]跟x的大小关系
	// x最左最右的情况
	s, e = max(0, s-k), min(s+k, len(arr)-1) // 往左右各扩展k个元素
	for e-s+1 > k {
		if arr[e]-x >= x-arr[s] {
			e--
		} else {
			s++
		}
	}
	return arr[s : e+1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
