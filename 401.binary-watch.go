 /*
 * @lc app=leetcode id=401 lang=golang
 *
 * [401] Binary Watch
 */
//  0 ms
// 回溯法， C(10, num)
func readBinaryWatch(num int) []string {
	rtn := []string{}
	for _,v := range generate(num) {
		rtn = append(rtn, time(v))
	}
	return rtn
}

func generate(num int) []int {
	rtn := []int{}
	dfs(0, num, 0, &rtn)
	return rtn
}

// s表示处理第几位，t剩余的1的个数
func dfs(s, t, cur int, ans *[]int){
	if s == 10 {
		if t == 0 && isValid(cur){
			(*ans) = append((*ans), cur)
		}
		return 
	}

	// 置0
	dfs(s+1, t, cur, ans)
	// 置1
	cur |= (1 << uint(s))
	dfs(s+1, t-1, cur, ans)
	// 此处不用再恢复cur 
}

func isValid(a int) bool{
	hour := (a&0x3c0) >> 6
	min := (a&0x3f)
	if hour>=0 && hour <=11 && min>=0 && min <=59 {
		return true
	}
	return false
}

// 10的灯，前4个为时，后6个为分
func time(num int) string {
	hour := (num&0x3c0) >> 6
	min := (num&0x3f)
	return fmt.Sprintf("%d:%02d", hour, min)
}

