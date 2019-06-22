/*
 * @lc app=leetcode id=841 lang=golang
 *
 * [841] Keys and Rooms
 */
//  bitmap标记优化空间？？？
// 典型错误思路：取钥匙的过程中，有的房间进不去，而遍历的时候进去了
func canVisitAllRooms(rooms [][]int) bool {
	l := len(rooms)
	roomMap := make([]bool, l)
	roomMap[0] = true // 一定注意，0肯定true，不管rooms是否存在
	for _,v := range rooms {
		for _,v1 := range v {
			roomMap[v1] = true
		}
	}

	for _,v := range roomMap {
		if v == false {
			return false
		}
	}
	return true
}

