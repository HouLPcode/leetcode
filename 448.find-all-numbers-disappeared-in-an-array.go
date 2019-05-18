/*
 * @lc app=leetcode id=448 lang=golang
 *
 * [448] Find All Numbers Disappeared in an Array
 */
 // 技巧：nums[nums[i] -1] = -nums[nums[i]-1]，没有变成负数的值即为没有出现的值
 // 把数组中出现的数字当成索引，将索引对应的num[index]取负数作为遍历过的标记
func findDisappearedNumbers(nums []int) []int {
	for i:= 0; i< len(nums); i++{
		if nums[i] < 0{
			if nums[(-nums[i])-1] < 0{// 已经访问过，直接跳过
				continue
			}
			nums[(-nums[i])-1] = -nums[(-nums[i])-1]
		}else{
			if nums[nums[i]-1] < 0{// 已经访问过，直接跳过
				continue
			}
			nums[nums[i]-1] = -nums[nums[i]-1]
		}
	}

	rtn := []int{}
	for k,v := range nums{
		if v > 0{
			rtn = append(rtn, k+1)//注意此处要的是索引k，不是数值v
		}
	}
	return rtn
}

