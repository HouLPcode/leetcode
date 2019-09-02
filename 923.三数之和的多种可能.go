/*
 * @lc app=leetcode.cn id=923 lang=golang
 *
 * [923] 三数之和的多种可能
 */
//  8 ms, faster than 75.00%
// [0,0,0]\n0
// [0,2,0]\n2
func threeSumMulti(A []int, target int) int {
	// map 统计每个数的频次
	times := make(map[int]int, 100)
	for _, v := range A {
		times[v]++
	}
	cnt := 0
	for k, v := range times {
		if k*3 == target { // AAA型
			cnt += v * (v - 1) * (v - 2) / 6
			cnt %= 1000000007
		}
		for k1, v1 := range times {
			if k1 <= k { // 去重
				continue
			}

			if k+k1+k1 == target { // ABB型
				cnt += v * (v1) * (v1 - 1) / 2
				cnt %= 1000000007
			}
			if k+k+k1 == target { // AAB型
				cnt += v * (v - 1) / 2 * v1
				cnt %= 1000000007
			}

			if target-k-k1 > k1 &&
				times[target-k-k1] > 0 { // ABC型
				cnt += v * v1 * times[target-k-k1]
				cnt %= 1000000007
			}
		}
	}
	return cnt
}
