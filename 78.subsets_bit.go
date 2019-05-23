/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */
 // 法1，回溯法
 // 法2，位运算，合理的值表示所有情况
 // [9,0,3,5,7] 回溯法的顺序不满足答案要求？？？
 func subsets(nums []int) [][]int {
	// 位运算，
	// sort.Ints(nums)
	count := 1 << uint(len(nums)) //总可能数 2^len(nums)
	result := make([][]int, 0, count)
	for i:=0; i<count; i++{
		path := make([]int, 0, len(nums)) //暂存单条结果
		for j:=0; j<len(nums); j++{
			if i & (1 << uint(j)) != 0{ // 注意此处是j位是1，不是数值1
				path = append(path, nums[j])
			}
		}
		data := make([]int, len(path))
		copy(data, path)
		sort.Ints(data)// 题目中没有要求此处必须排序
		result = append(result, data)
	}
	return result
}
