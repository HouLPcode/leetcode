/*
 * @lc app=leetcode id=473 lang=golang
 *
 * [473] Matchsticks to Square
 */
 // 位运算处理
//  计算正方形边长，找到所有能组合成该长度的集合，然后挑出4个子集，无重复元素，且元素个数为nums，则成功构成正方形
// [2,2,2,2,2,6] 
func makesquare(nums []int) bool {
	if len(nums) < 4{
		return false
	}
	sum := 0
	for i:=0; i<len(nums); i++ {
		sum += nums[i]
	}
	if sum % 4 != 0{
		return false
	}
	target := sum / 4
	// 题目要求：nums长度不超过15，长度和不超过10^9
	sides := make([]int, 0)
	for i:=0; i< 1<<uint(len(nums)); i++{ 		//找出i中所有1的位置
		total := 0
		for j:=0; j<len(nums); j++{		// 注意此处的技巧，遍历nums，在i中找位1
			if i & (1<<uint(j)) != 0{ // i的第j位为1，即i中有nums[j],且i中所有位1组成合法边，这样不用单独保存需要的边具体是怎么组成的
				total += nums[j]
			}
			if total == target{
				sides = append(sides, i) // 此处只需要保存i即可，不需要保存具体的nums[x]
			}
		}	
	}
	//之后在size中选出4个没有重复元素，且所有元素都用上的组合，能找到返回true. 类似4sum？？？
	// 1. 先找两两不想交的
	sides_2 := make([]int, 0, len(sides)/2)
	for i:=0; i<len(sides); i++ { // O(n^2),可以优化吗？？？
		for j:=i+1; j<len(sides); j++ {
			if sides[i] & sides[j] == 0{
				sides_2 = append(sides_2, sides[i]|sides[j])//此处合并两个边，避免使用二维数组
			}
		}
	}
	// 2. 找四个都不想交的
	for i:=0; i<len(sides_2); i++ {
		for j:=i+1; j<len(sides_2); j++ {
			if sides_2[i] & sides_2[j] == 0 { //因为i和j的周长是定值，所以此处合一保证所有的数字都用上了。如果少数字，不可能是4边两两不想交
				return true
			}
		}
	}
	return false
}


