/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
func twoSum(nums []int, target int) []int {
	// map 存储值
	cnt := []int{}
	nummap := make(map[int]int,0)
	// nums中的数值为k，位置为val
	// 没必要首先遍历完所有数据，浪费时间和空间
	for k,v := range nums{
		nummap[v] = k
	}

	for k,v := range nums{
		// 注意，数值不能重复使用 [3,2,4] 6
		if val,ok := nummap[target - v]; ok && k != val{
			cnt = append(cnt,k)
			cnt = append(cnt,val)
			break
		} 
	}
	return cnt
}

