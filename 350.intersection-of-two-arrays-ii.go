/*
 * @lc app=leetcode id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 */
//  4 ms 95.38 %
//  [4,9,5]\n[9,4,9,8,4]
// [3,1,2]\n[1,1]
// map 处理
func intersect(nums1 []int, nums2 []int) []int {
	num1map := make(map[int]int, 0)
	for _,v := range nums1 {
		num1map[v]++
	}
	num2map := make(map[int]int, 0)
	for _,v := range nums2 {
		num2map[v]++
	}
	rtn := make([]int, 0, len(nums1))
	for k,v := range num1map {
		for i:=0; i < v && i < num2map[k]; i++ { //注意此处的处理，通过比较两个map的次数来添加最后的结果中次数
			rtn = append(rtn, k)
		}
	}	
	return rtn
}

