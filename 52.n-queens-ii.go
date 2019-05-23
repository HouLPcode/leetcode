/*
 * @lc app=leetcode id=52 lang=golang
 *
 * [52] N-Queens II
 */
func totalNQueens(n int) int {
	result := 0
	mark := make([][]bool, n) //标记数组，表示放置皇后的影响范围
	for k := range mark{
		mark[k] = make([]bool, n)
	}

	generate(0, n, &mark, &result)
	return result    
}


func generate(k, n int, mark *[][]bool, result *int){
	// 放第k行的皇后，范围 0...n-1
	if k == n{// 已经放置完了，记录这个解，是一个合法解
		(*result)++
		return
	}
	
	for i:=0; i<n; i++{// 在第k行找一个位置放置皇后
		if (*mark)[k][i] == false{ // i列上可以放置
			markCopy := make([][]bool, len(*mark))// 注意此处二维slice的备份方式
			for i:=0; i<len(markCopy); i++{
				markCopy[i] = make([]bool, len((*mark)[0]))
				copy(markCopy[i], (*mark)[i])
			}
			putQueue(k,i,mark)// 放置后辐射影响范围
			generate(k+1,n,mark,result)//递归下一行放置
			// 处理剩余可能的i的位置，需要恢复i辐射前的mark,恢复之前的location
			for i:=0; i<len(markCopy); i++{
				copy((*mark)[i], markCopy[i])
			}
		}
	}
}

// 方向数组
var dx []int = []int{-1, 1, 0, 0, -1, 1, -1, 1}
var dy []int = []int{0, 0, -1, 1, -1, 1, 1, -1}

// 在x,y放置皇后,影响的范围
func putQueue(x, y int, mark *[][]bool){
	(*mark)[x][y] = true
	for i:=1; i<len(*mark); i++{//扩展次数1...n-1
		for j:=0; j<8; j++{// 8个方向
			curX := x+i*dx[j]
			curY := y+i*dy[j]
			if 0<=curX && curX<len(*mark) &&
			   0<=curY && curY<len((*mark)[0]){
				 (*mark)[curX][curY] = true  
			   }
		}
	}
}



