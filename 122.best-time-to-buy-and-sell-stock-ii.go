/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */
func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0{
		return 0
	}
	// DP方法思路分析
	// 假设f(i) 表示第i次的最大利润
	// 最大利润的时候手里肯定没有股票，所以f(i-1)的时候手里面也没有股票，没法卖股票，i的时候只能什么也不做，有贪心的思想，无法获取i时候的利润
	// 所以做出以下定义，f(i)同时记录有股票和没股票时候的最大利润
	
	// 合理的定义
	//f[i][0]第i次没有股票是的最大利润，f[i][1]第i次有股票的最大利润
	// 买卖次数无限制
	// 转移方程 f[i][0] = max( f[i-1][0],f[i-1][1]+prices[i])
	// 		   f[i][1] = max(f[i-1][0]-prices[i],f[i-1][1])
	// 最终结果f[i][0]
	buf := make([][2]int,len(prices),len(prices))
	buf[0][1] = -prices[0]

	for i:=1; i<len(prices); i++{
		buf[i][0] = max(buf[i-1][0],buf[i-1][1]+prices[i])
		buf[i][1] = max(buf[i-1][0]-prices[i],buf[i-1][1])
	}
	return buf[len(prices)-1][0]
}

func max(a, b int)int{
	if a > b{
		return a
	}
	return b
}

