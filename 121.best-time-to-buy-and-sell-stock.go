/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

 func maxProfit(prices []int) int {
    if len(prices) == 0{// 需要处理空情况
        return 0
    }
    
    // 遍历一遍，记录当前已知最小值，计算当前值-最小值的收益，统计最大的收益
    curMin,profit := prices[0], 0
	for _,val := range prices{
		curMin = min(val,curMin)
		// 不会出现先卖后买的情况
		profit = max(profit,val - curMin)
	}
	return profit
}
func max(a, b int)int{
	if a > b{
		return a
	}
	return b
}

func min(a, b int)int{
	if a < b{
		return a
	}
	return b
}

