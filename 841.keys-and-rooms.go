/*
 * @lc app=leetcode id=841 lang=golang
 *
 * [841] Keys and Rooms
 */
//  4 ms 100 %
//  bitmap标记优化空间？？？
// DFS 实现
func canVisitAllRooms(rooms [][]int) bool {
	l := len(rooms)
	roomMap := make([]bool, l)	
	// dfs 尝试
	dfs(rooms, 0, roomMap)
	for _,v := range roomMap {
		if v == false {
			return false
		}
	}
	return true
}
// 访问第s房间
func dfs(rooms [][]int, s int, roomMap []bool) {
	if roomMap[s] == true {
		return //访问过，直接返回
	}
	roomMap[s] = true
	for _,v := range rooms[s] {
		dfs(rooms, v, roomMap)
	}
}

