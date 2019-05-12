/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 */
func isHappy(n int) bool {
	happy := make(map[int]bool)
	for {
		if _,ok := happy[n]; ok{
			return false // 再次访问表示出现循环
		}
		happy[n] = true 
		
		n = calc(n)
		if n == 1{
			return true
		}
	}
	return false
}

func calc(num int)int{
	cnt := 0
	for num != 0{
		cnt += (num%10)*(num%10)
		num /= 10 
	}
	return cnt
}

