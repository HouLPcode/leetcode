/*
 * @lc app=leetcode id=987 lang=golang
 *
 * [987] Vertical Order Traversal of a Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // 0ms
 func verticalTraversal(root *TreeNode) [][]int {
	// 缓存每个节点的坐标， []int{x, y, val}
	pos := [][3]int{}
	getPos(root, 0, 0, &pos)
	//按照 x,y,val的顺序排序
	sort.Slice(pos, func (i, j int)bool  {
		if pos[i][0] > pos[j][0] {
			return false
		} else if pos[i][0] == pos[j][0] &&
			pos[i][1] < pos[j][1]{ // 注意，y的顺序是由大到小
				return false
		} else if pos[i][0] == pos[j][0] &&
		pos[i][1] == pos[j][1] &&
		pos[i][2] > pos[j][2] {
			return false
		}
		return true
	})

	// 输出时候，每个x为一组
	rtn := [][]int{}
	rtn = append(rtn, []int{pos[0][2]})
	curX := pos[0][0]
	for i:=1; i<len(pos); i++ {
		if pos[i][0] == curX{
			rtn[len(rtn)-1] = append(rtn[len(rtn)-1], pos[i][2])
		} else {
			curX = pos[i][0]
			rtn = append(rtn, []int{pos[i][2]})
		}
	}
	return rtn
}

// pre遍历， 也可以说是dfs
func getPos(root *TreeNode, x, y int, pos *[][3]int) {
	if root == nil {
		return 
	}
	*pos = append(*pos, [3]int{x, y, root.Val})
	getPos(root.Left, x-1, y-1, pos)
	getPos(root.Right, x+1, y-1, pos)
}
