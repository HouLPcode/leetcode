/*
 * @lc app=leetcode id=941 lang=golang
 *
 * [941] Valid Mountain Array
 */
//  24 ms 96.72 %
//  [9,8,7,6,5,4,3,2,1,0]
func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}
	i := 0
	for i+1 < len(A) && A[i] < A[i+1] {
		i++
	} // i为山顶
	if i == len(A)-1 || i == 0 { //此处一定注意需要判断i是否0或len-1
		return false // 递增或递减
	}
	for i+1 < len(A) && A[i] > A[i+1] {
		i++
	}
	if i == len(A)-1 {
		return true
	}
	return false
}

