/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */
 // 法1，回溯法
 // 法2，位运算，合理的值表示所有情况
 // [9,0,3,5,7] 回溯法的顺序不满足答案要求？？？
//  ✘ answer:          [[],[4],[1],[0],[1,4],[0,1],[0,1],[0,1,4]]
//  ✘ expected_answer: [[],[4],[1],[1,4],[0],[0,4],[0,1],[0,1,4]]
 
func subsets(nums []int) [][]int {
	result := [][]int{}
	for i:=0; i<=len(nums); i++{
		generate(0, i,[]int{}, nums, &result)
	}
	return result
}

func generate(s int,n int, path []int, nums []int, result *[][]int){
	if n == len(path){
		data := make([]int, len(path))
		copy(data, path)
		sort.Ints(data)
		*result = append((*result), data)
		return 
	}

	for i:=s; i<len(nums); i++ {
		path = append(path, nums[i])
		generate(i+1, n, path, nums, result)
		path = path[:len(path)-1]
	}
}

