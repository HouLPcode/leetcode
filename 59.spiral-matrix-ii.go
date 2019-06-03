/*
 * @lc app=leetcode id=59 lang=golang
 *
 * [59] Spiral Matrix II
 */
func generateMatrix(n int) [][]int {
	rtn := make([][]int,n)
	for i:=0; i<n; i++ {
		rtn[i] = make([]int, n)
	}

	start := 1
	for i:=0; i<(n+1)/2; i++ {
		circle(i,i,n-1-i,n-1-i,&rtn,&start)
	}

	return rtn
}

// 左上和右下坐标， n为计数起始值
func circle(sr, sc, er, ec int, matrix *[][]int, n *int) {
	// 上 [,]
	for c:=sc; c<=ec; c++ {
		(*matrix)[sr][c] = (*n)
		(*n)++
	}
	// 右 (,]
	for r:=sr+1; r<=er; r++ {
		(*matrix)[r][ec] = (*n)
		(*n)++
	}

	if sr != er && sc != ec { // 不是一行且不是一列，该题目中判断不是必须的，因为是正方形
		// if  下 (,]
		for c:=ec-1; c >= sc; c-- {
			(*matrix)[er][c] = (*n)
			(*n)++
		}
		// 左 (,)
		for r:=er-1; r > sr; r-- {
			(*matrix)[r][sc] = (*n)
			(*n)++
		}
	}
}

