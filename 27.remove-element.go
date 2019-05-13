/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */
// 测试用例，[4,5] 4
//   nums = [3,2,2,3], val = 3   
//   nums = [0,1,2,2,3,0,4,2], val = 2 
func removeElement(nums []int, val int) int {
	// 两个首尾指针，逐步遍历
	// 首指针遇到val值时和尾指针交换，尾指针遇到val时直接跳过
	i, j := 0, len(nums)-1
	if j == -1{ // 有必要吗？？？
		return 0
	}
	cnt := 0
	for i != j{
		if nums[i] == val && nums[j] == val{
			j--
		}else if nums[i] == val && nums[j] != val{
			// cnt += 1
			nums[i], nums[j] = nums[j], nums[i]
			// j--
		}else if nums[i] != val && nums[j] == val{
			j--
		}else{
			i++
			cnt++ // 只在这一种情况下cnt++
		}
	}
	if nums[i] != val{//可能会少计算1个 i == j 时候
		cnt++
	}
	return cnt
}

