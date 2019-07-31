import "sort"

/*
 * @lc app=leetcode id=853 lang=golang
 *
 * [853] Car Fleet
 */
// (32 ms) √ Your runtime beats 89.66 %
//   × testcase: '10\n[2,4]\n[3,2]'
//   × testcase: '10\n[8,3,7,4,6,5]\n[4,4,4,4,4,4]'

// 注意不能超车
// 思路：先按照起始位置position从大到小排序，然后计算到达终点的时间，后面时间短的表示超车前面车辆，属于一个车队
// target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]
func carFleet(target int, position []int, speed []int) int {
	cars := make([]Car, len(position))

	// 计算出到达终点需要的时间
	for i := 0; i < len(speed); i++ {
		cars[i].Position = position[i]
		cars[i].Time = float64((target - position[i])) / float64(speed[i]) // 注意时间采用的是浮点数，整数会出错
	}

	// 根据起始位置对时间排序
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].Position > cars[j].Position
	})

	// 统计车队数量，注意不能超车
	cnt := 0
	curMax := 0.0 // 耗时>=0
	for i := 0; i < len(cars); {
		curMax = cars[i].Time
		cnt++
		for i < len(cars) && curMax >= cars[i].Time {
			i++
		}
	}
	return cnt
}

type Car struct {
	Time     float64
	Position int
}
