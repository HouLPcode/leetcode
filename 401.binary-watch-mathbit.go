 /*
 * @lc app=leetcode id=401 lang=golang
 *
 * [401] Binary Watch
 */
//  0 ms
//  数学方式， bit操作构成组合数， 验证所有可能
func readBinaryWatch(num int) []string {
	rtn := make([]string, 0) // 一共C(10, num)组合个可能
	for i:=0; i<(1<<10); i++ {
		if isValid(i, num) == true {
			rtn = append(rtn, time(i))
		}
	}
	return rtn
}

func isValid(a, num int) bool{
	hour := (a&0x3c0) >> 6
	min := (a&0x3f)
	if hour>=0 && hour <=11 && min>=0 && min <=59 {
		// 1的个数
		cnt := 0
		for a != 0 {
			cnt++
			a &= (a-1)
		}
		return cnt == num
	}
	return false
}

// 10的灯，前4个为时，后6个为分
func time(num int) string {
	hour := (num&0x3c0) >> 6
	min := (num&0x3f)
	return fmt.Sprintf("%d:%02d", hour, min)
}

