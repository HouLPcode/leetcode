/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */
 
 //贪心不好使 [1,6,7] 30 
 //          6 * 5  => 5
 //          7 * 4 + 1 * 2 => 6
func coinChange(coins []int, amount int) int {
	// f(i) amount为i的最少硬币数
	// f(i) = 1 + min( f(i-coins[...]) )
	// 最终结果即为 f(amount)
	
	// 怎么确认最终没有解决方案，返回 -1
	// 初始值不能设置成-1，否则min函数一直返回-1
	// 由于硬币面额均为正数，所以硬币个数不可能超过总金额cmount，故将所有值初始化为amount+1,表示没有访问过
	buf := make([]int,amount+1,amount+1)//0的存在，创建的时候注意要多1个
	for i := range buf{
		buf[i] = amount + 1
	}
	buf[0] = 0 //这一句很重要，0个硬币组成金额0，且金额0不会访问buf数组
	for i:=0; i<=amount; i++{//最后需要的是amount，此处是小于等于
		for _,coin := range coins{
			if i >= coin{
				buf[i] = min(buf[i],buf[i - coin]+1)
			}
		}
	}
	if buf[amount] == amount + 1{ //数组没有访问过，即无法组成该金额
		return -1
	}
	return buf[amount]
}

func min(a,b int)int{
	if a < b {
		return a
	}
	return b
}

