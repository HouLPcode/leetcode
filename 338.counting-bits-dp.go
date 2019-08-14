/*
 * @lc app=leetcode id=338 lang=golang
 *
 * [338] Counting Bits
 *
 * https://leetcode.com/problems/counting-bits/description/
 *
 * algorithms
 * Medium (64.44%)
 * Total Accepted:    162.6K
 * Total Submissions: 252.3K
 * Testcase Example:  '2'
 *
 * Given a non negative integer number num. For every numbers i in the range 0
 * ≤ i ≤ num calculate the number of 1's in their binary representation and
 * return them as an array.
 *
 * Example 1:
 *
 *
 * Input: 2
 * Output: [0,1,1]
 *
 * Example 2:
 *
 *
 * Input: 5
 * Output: [0,1,1,2,1,2]
 *
 *
 * Follow up:
 *
 *
	* It is very easy to come up with a solution with run time
	* O(n*sizeof(integer)). But can you do it in linear time O(n) /possibly in a
	* single pass?
 * Space complexity should be O(n).
 * Can you do it like a boss? Do it without using any builtin function like
 * __builtin_popcount in c++ or in any other language.
 *
*/
// 840 ms, faster than 74.35% 不用优化
// 类似dp求解
// dp[i] 表示 数字i中1的个数
// dp[i] = dpp[i&(i-1)] + 1 ,最右边的1清零后的个数+1
// dp[0] = 0
// dp[1] = dp[0] + 1
// dp[2] = dp[0] + 1
// dp[3] = dp[2] + 1
// dp[4] = dp[0] + 1
func countBits(num int) []int {
	rtn := make([]int, num+1)
	for i := 1; i <= num; i++ {
		rtn[i] = rtn[i&(i-1)] + 1
	}
	return rtn
}
