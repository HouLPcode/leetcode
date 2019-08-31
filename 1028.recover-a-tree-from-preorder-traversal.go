/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 12 ms, faster than 95.92%
func recoverFromPreorder(S string) *TreeNode {
	nodes := [][]int{} // [depth,val]
	for i := 0; i < len(S); {
		dep := 0
		for i < len(S) && S[i] == '-' {
			dep++
			i++
		}
		node := []int{dep, 0}
		l := i
		for i < len(S) && S[i] != '-' {
			i++
		}
		node[1], _ = strconv.Atoi(S[l:i])
		nodes = append(nodes, node)
	}
	return pre(nodes)
}

//
func pre(nodes [][]int) *TreeNode {
	if len(nodes) == 0 {
		return nil
	}
	if len(nodes) == 1 {
		return &TreeNode{nodes[0][1], nil, nil}
	}

	root := &TreeNode{}
	root.Val = nodes[0][1]
	dep := nodes[0][0]

	if nodes[1][0] == nodes[0][0]+1 {
		// 找 dep+1 切分两块
		i := 2
		for ; i < len(nodes); i++ {
			if nodes[i][0] == dep+1 {
				break
			}
		}
		if i == len(nodes) {
			// 右子树为nil
			root.Right = nil
		} else {
			root.Right = pre(nodes[i:])
		}
		root.Left = pre(nodes[1:i])
	} else {
		// 左子树为nil
		root.Left = nil
		// 找 dep+1 切分两块
		i := 2
		for ; i < len(nodes); i++ {
			if nodes[i][0] == dep+1 {
				break
			}
		}
		if i == len(nodes) {
			// 右子树为nil
			root.Right = nil
		} else {
			root.Right = pre(nodes[i:])
		}
	}

	return root
}
