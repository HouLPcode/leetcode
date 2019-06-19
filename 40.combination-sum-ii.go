/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */
//  0 ms
// [3,1,3,5,1,1]\n8
 // 输入数组有重复元素，每个元素只能用一次
 // 难点：怎么舍去重复解  先对candidates排序，之后在一层中，不是第一个且出现过就直接忽略这个解
func combinationSum2(candidates []int, target int) [][]int {
	rtn := [][]int{}
	sort.Ints(candidates) // 注意此处排序后在后面进行剪枝，节省时间
	dfs(candidates, 0, target, []int{}, &rtn)
	return rtn
}

func dfs(candidates []int, s, target int, curr []int, ans *[][]int){
	if target == 0 {
		// 一定注意，此处深层拷贝，直接使用curr会出错
		data := make([]int, len(curr))
		copy(data, curr)
		(*ans) = append((*ans), data)
		return
	}
	for i:=s; i<len(candidates); i++ {
		if candidates[i] > target {
			return // 剪枝，注意此处是return不是break，表示i及其后面的多有带选项都不满足条件，直接放弃
		}
		if i>s && candidates[i] == candidates[i-1] { //不是首次出现，且之前在该层出现过，则直接忽略
			continue
		}
		
		curr = append(curr, candidates[i])
		dfs(candidates, i+1, target-candidates[i], curr, ans) // 每次调用进入新的一层
		curr = curr[:len(curr)-1]
	}
}
