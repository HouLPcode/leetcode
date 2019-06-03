/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */
 // 法1，回溯法
 // 法2，位运算，合理的值表示所有情况
 // [9,0,3,5,7] 回溯法的顺序不满足答案要求？？？

func subsets(nums []int) [][]int {
	result := [][]int{}
	generate(0, []int{}, nums, &result)
	return result
}

func generate(s int, path []int, nums []int, result *[][]int){
	data := make([]int, len(path))
	copy(data, path)
	sort.Ints(data)
	*result = append((*result), data)

	for i:=s; i<len(nums); i++ {
		path = append(path, nums[i])
		generate(i+1, path, nums, result)
		path = path[:len(path)-1]
	}
}
