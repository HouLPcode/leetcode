/*
 * @lc app=leetcode id=733 lang=golang
 *
 * [733] Flood Fill
 */
 // [[0,0,0],[0,1,1]]\n1\n1\n1  新旧颜色一致，导致超时
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	// 题目中给出长宽范围[1, 50]
	// DFS方式
	floodFillDFS(&image, sr, sc, newColor)
	return image
}

var dx = []int{1, -1, 0,  0}
var dy = []int{0,  0, 1, -1}

func floodFillDFS(image *[][]int, r, c, newColor int){
	curColor := (*image)[r][c]
	if curColor == newColor{ //新旧颜色一致，避免死循环
		return 
	}
	(*image)[r][c] = newColor //染色

	for i:=0; i<4; i++{ // 扩散
		if 0<=r+dx[i] && r+dx[i]<=len((*image))-1 &&
		   0<=c+dy[i] && c+dy[i]<=len((*image)[0])-1 &&
		   (*image)[r+dx[i]][c+dy[i]] == curColor {//len指针？？
			floodFillDFS(image, r+dx[i], c+dy[i], newColor)
		}
	}
}
