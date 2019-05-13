/*
 * @lc app=leetcode id=167 lang=golang
 *
 * [167] Two Sum II - Input array is sorted
 */
func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	//首尾逐步排查 O(n)
	// 二分法可以吗？？？
	for i < j { // == ????
		if numbers[i] + numbers[j] == target{
			break
		}else if numbers[i] + numbers[j] < target{
			i++
		}else{
			j--
		}
	}
	return []int{i+1,j+1}// 题目要求肯定有解
}

