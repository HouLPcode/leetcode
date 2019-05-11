/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 */
func fourSum(nums []int, target int) [][]int {
    if len(nums) < 4{
		return [][]int{}
	}
	rtn := [][]int{}
	sort.Ints(nums)
	for a := 0; a < len(nums)-3; a++{
		if a > 0 && nums[a] == nums[a-1]{
			continue
		}
		for b:=a+1; b < len(nums)-2; b++{
			if b > a+1 && nums[b] == nums[b-1]{//去重
				continue
			}
			for c, d := b+1,len(nums)-1; c < d;{
				if nums[a] + nums[b] + nums[c] + nums[d] == target{
					rtn = append(rtn,[]int{nums[a],nums[b],nums[c],nums[d]})
					// 找到后，后续可能还有与a匹配的结果，继续寻找
					for c++; nums[c] == nums[c-1] && c < d; c++ {} //无限的++可能导致nums越界
					for d--; nums[d] == nums[d+1] && c < d; d-- {}
				}else if nums[a] + nums[b] + nums[c] + nums[d] < target{
					c++
				}else if nums[a] + nums[b] + nums[c] + nums[d] > target{
					d--
				}
			}
		}
	}
	return rtn
}

