/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
func twoSum(nums []int, target int) []int {
	// map 存储值
	cnt := []int{}
	nummap := make(map[int]int,0)
	for k,v := range nums{
		nummap[v] = k
	}

	for k,v := range nums{
		// 注意，数值不能重复使用 [3,2,4] 6
		if val,ok := nummap[target - v];ok{
			cnt = append(cnt,k)
			cnt = append(cnt,val)
			break
		} 
	}
	return cnt
}

