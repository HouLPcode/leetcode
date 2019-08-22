/*
 * @lc app=leetcode id=1089 lang=golang
 *
 * [1089] Duplicate Zeros
 */
//  (28 ms) √ Your runtime beats 99.55 %
// [0,0,0,0,0,0,0]
func duplicateZeros(arr []int) {
	// 遍历一遍统计0的个数
	cnt0 := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			cnt0++
		}
	}

	// 填充了cnt0个0，从右往左过滤掉 cnt个数值，其中0站2个
	j := len(arr) - 1
	for cnt0 > 0 { // 有可能把c变成-1
		if arr[j] == 0 {
			cnt0 -= 2
		} else {
			cnt0--
		}
		j--
	}

	i := len(arr) - 1
	if cnt0 == -1 { // 留下一个0处理, 一定注意次数，可能存在某个0填充时，正好填充到末尾，空间不足，新加的0没有加进去
		arr[i] = 0
		i--
	}

	for i >= 0 {
		arr[i] = arr[j]
		i--
		if arr[j] == 0 { // 填充一个0
			arr[i] = 0
			i--
		}
		j--
	}
}
