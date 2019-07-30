import "math"

/*
 * @lc app=leetcode id=458 lang=golang
 *
 * [458] Poor Pigs
 */
//  0 ms
// 每15分钟喝一次，一小时能喝4次。一定注意，四段有5个端点。
// 喝水时间    0  15  30  45
// 出结果时间  15 30  45  60
// 所以一直猪可以喝4个桶，所以可以区分5个桶，喝四个剩1个
// 两只猪可以区分 5^2个。
// 错误思路：先平分5组，一直猪喝，能确认分组，然后在分组中确认那个有毒
// 正确思路：两只猪同时喝水，把25分成5*5矩阵，A猪按行喝水，B猪按列喝水，最后取交集
// 公式： (p/m + 1) ^ ans >= buckets
func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	// ans >= log(buckets) / log(p/m+1)
	ans := math.Log(float64(buckets)) / math.Log(float64(minutesToTest/minutesToDie+1))
	if int(math.Pow(float64(minutesToTest/minutesToDie+1), float64(int(ans)))) < buckets {
		return int(ans) + 1
	}
	return int(ans)
}
