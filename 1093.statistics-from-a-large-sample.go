/*
 * @lc app=leetcode id=1093 lang=golang
 *
 * [1093] Statistics from a Large Sample
 */
// 0 ms
func sampleStats(count []int) []float64 {
	rtn := make([]float64, 5)
	sum := 0
	cnt := 0
	cntMax := 0
	for k, v := range count {
		if v == 0 {
			continue
		}
		// if cntrtn[0] == 0.0 { // 最小值有可能是0,该条件判断最小值错误
		if cnt == 0 {
			rtn[0] = float64(k) // min, 仅第一次进入时赋值一次
		}
		rtn[1] = float64(k) // max， 每次都要赋值
		sum += k * v
		cnt += v
		// 众数
		if v > cntMax {
			rtn[4] = float64(k)
			cntMax = v
		}
	}
	// 平均数
	rtn[2] = float64(sum) / float64(cnt)
	// 中位数
	if cnt&1 == 0 { // 偶数
		c := 0
		for k, v := range count {
			if c+v < cnt/2 {
				c += v
			} else if c+v == cnt/2 {
				i := k + 1
				for ; count[i] == 0; i++ {
				}
				rtn[3] = float64(k+i) / 2.0
				break
			} else {
				rtn[3] = float64(k)
				break
			}
		}
	} else { // 奇数
		c := 0
		for k, v := range count {
			if c+v < cnt/2 { // 注意，先+个数，再判断
				c += v
			} else {
				rtn[3] = float64(k)
				break
			}
		}
	}
	return rtn
}
