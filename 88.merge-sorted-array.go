/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */
func merge(nums1 []int, m int, nums2 []int, n int)  {
	// 从后往前遍历，避免数组元素的多次移动

	// for m--,n--; m >=0 && n >= 0;{//语法错误
	for m, n = m-1, n-1; m >=0 && n >= 0; {
		if nums1[m] > nums2[n]{
			nums1[m+n+1] = nums1[m]
			m--
		}else{
			nums1[m+n+1] = nums2[n]
			n--
		}
	}
	//注意收尾工作，nums1处理完，nums2还没有处理完，将nums2 拷贝到nums1
	for ;n >= 0; n--{
		nums1[n] = nums2[n]
	}
}
 // TODO 待优化
// Your runtime beats 19.11 % of golang submissions
// Your memory usage beats 5.56 % of golang submissions (3.6 MB)

