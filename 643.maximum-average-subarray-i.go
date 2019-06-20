/*
 * @lc app=leetcode id=643 lang=golang
 *
 * [643] Maximum Average Subarray I
 */
//  124 ms 97.22 %
func findMaxAverage(nums []int, k int) float64 {
    if len(nums) < k {
		return 0.0
	}
	sum := 0
	for i:=0; i<k; i++ {
		sum += nums[i]
	}
	sumMax := sum
	for i:=k; i<len(nums); i++ {
		sum += nums[i]
		sum -= nums[i-k]
		if sum > sumMax {
			sumMax = sum
		}
	}
	return float64(sumMax) / float64(k)
}

