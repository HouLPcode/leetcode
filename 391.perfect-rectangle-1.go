// 56 ms, faster than 96.77%
// 限制条件 ：
// 1. 统计每个方块的四个角出现的次数，一定注意，可能次数不不不不不是 1，2，4，可能有3的情况
// 2. 组合后的大面积与小面积相同（注意，不能只看这个条件，比如覆盖面积和缺失面积相同的时候也满足这个条件）
// 3. 次数为1的四个点是最大最小的四个点
// [[0,-1,1,0],[0,0,1,1],[0,1,1,2],[0,2,1,3]]
// [[1,1,3,3],[3,1,4,2],[3,2,4,4],[1,3,2,4],[2,3,3,4]]
// [[1,1,2,2],[1,1,2,2],[2,1,3,2]]
// [[0,0,3,3],[1,1,2,2],[1,1,2,2]]
// [[0,0,1,1],[0,0,2,1],[1,0,2,1],[0,2,2,3]]
// [[0,0,1,1],[0,1,1,2],[0,2,1,3],[1,0,2,1],[1,0,2,1],[1,2,2,3],[2,0,3,1],[2,1,3,2],[2,2,3,3]]
func isRectangleCover(rectangles [][]int) bool {
	//map 记录各个角出现的次数,因为是偶数次，保留奇数次，所以bool可行
	minmaxAngles := getAngles(rectangles[0]) // 记录四个角
	// 注意：此处不用string当作key，而是 x<<16+y  (int) 注意，x<<16|y坐标可能是负值，导致出错
	numMap := make(map[int]int, 0)
	sumArea := 0
	for _, v := range rectangles {
		sumArea += getArea(v)
		a := getAngles(v)
		numMap[a[0][0]<<16+a[0][1]]++
		numMap[a[1][0]<<16+a[1][1]]++
		numMap[a[2][0]<<16+a[2][1]]++
		numMap[a[3][0]<<16+a[3][1]]++

		if a[0][0] <= minmaxAngles[0][0] && a[0][1] <= minmaxAngles[0][1] { // ld
			minmaxAngles[0][0] = a[0][0]
			minmaxAngles[0][1] = a[0][1]
		}
		if a[1][0] >= minmaxAngles[1][0] && a[1][1] >= minmaxAngles[1][1] { // rt
			minmaxAngles[1][0] = a[1][0]
			minmaxAngles[1][1] = a[1][1]
		}
		if a[2][0] <= minmaxAngles[2][0] && a[2][1] >= minmaxAngles[2][1] { // lt
			minmaxAngles[2][0] = a[2][0]
			minmaxAngles[2][1] = a[2][1]
		}
		if a[3][0] >= minmaxAngles[3][0] && a[3][1] <= minmaxAngles[3][1] { // rd
			minmaxAngles[3][0] = a[3][0]
			minmaxAngles[3][1] = a[3][1]
		}
	}

	cnt := 0
	for k, v := range numMap {
		if v == 1 {
			cnt++
			if k != minmaxAngles[0][0]<<16+minmaxAngles[0][1] &&
				k != minmaxAngles[1][0]<<16+minmaxAngles[1][1] &&
				k != minmaxAngles[2][0]<<16+minmaxAngles[2][1] &&
				k != minmaxAngles[3][0]<<16+minmaxAngles[3][1] {
				return false
			}
		} else if v&1 != 0 {
			return false
		}
	}
	if cnt != 4 {
		return false
	}
	if sumArea != (minmaxAngles[1][0]-minmaxAngles[0][0])*(minmaxAngles[1][1]-minmaxAngles[0][1]) {
		return false
	}
	return true
}

// 计算面积
func getArea(a []int) int {
	return (a[2] - a[0]) * (a[3] - a[1])
}

// 计算几个角
func getAngles(a []int) [][]int { // 角的坐标值可能超过byte
	return [][]int{
		[]int{a[0], a[1]}, //ld
		[]int{a[2], a[3]}, //rt
		[]int{a[0], a[3]}, //lt
		[]int{a[2], a[1]}, //rd
	}
}

// func getAngles(a []int) [][]byte{// 角的坐标值可能超过byte
//     return [][]byte{
//         []byte{byte(a[0])+'0', byte(a[1])+'0'},//ld
//         []byte{byte(a[2])+'0', byte(a[3])+'0'},//rt
//         []byte{byte(a[0])+'0', byte(a[3])+'0'},//lt
//         []byte{byte(a[2])+'0', byte(a[1])+'0'},//rd
//     }
// }
