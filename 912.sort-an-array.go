/*
 * @lc app=leetcode id=912 lang=golang
 *
 * [912] Sort an Array
 */
//   880 ms 48.30%
// merge sort 归并排序
func sortArray(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}
	return merge(sortArray(nums[:l/2]), sortArray(nums[l/2:]))
}

func merge(nums1, nums2 []int) []int {
	rtn := make([]int, 0, len(nums1)+len(nums2))
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			rtn = append(rtn, nums1[i])
			i++
		} else {
			rtn = append(rtn, nums2[j])
			j++
		}
	}
	// 一定注意，此处推出循环时i或j某一个已经超过范围，一定不要写成nums[i+1:]或nums[j+1:]
	rtn = append(rtn, nums1[i:]...)
	rtn = append(rtn, nums2[j:]...)
	return rtn
}
