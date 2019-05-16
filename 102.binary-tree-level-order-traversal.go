/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归实现 
// 测试集 [1,2,3,4,5]
func levelOrder(root *TreeNode) [][]int {
	rnt := [][]int{} //零值,len外层是0
	relevel(root, 1, &rnt) //从1层开始计数，没有0层
	return rnt
}

func relevel(root *TreeNode, level int, result *[][]int){
	if root == nil{
		return 
	}
	depth := len((*result))
	if  depth < level{		// 该root节点是新的一层
		(*result) = append((*result),[]int{root.Val})
	}else{
		// depth该语句会导致每次节点都插入到最后一层，此处应该插入在level层
		// (*result)[depth-1] = append((*result)[depth-1],root.Val)
		(*result)[level-1] = append((*result)[level-1],root.Val)
	}
	relevel(root.Left, level+1, result)
	relevel(root.Right, level+1, result)
}

// func levelOrder(root *TreeNode) [][]int {
// 	//怎么区分一层,
// 	// 一次处理一层,每层生成一个[]slice，之后其添加到[][]
// 	if root == nil{
// 		return [][]int{}
// 	}
// 	out := [][]int{}
// 	list := list.New()
// 	list.PushBack(root)
// 	for list.Front()!=nil{
// 		levelLen := list.Len()
// 		currentLevel := []int{}
// 		for i :=0;i<levelLen;i++{
// 			node := list.Front()
// 			list.Remove(node)
// 			currentLevel = append(currentLevel,node.Value.(*TreeNode).Val)
// 			if node.Value.(*TreeNode).Left != nil{
// 				list.PushBack(node.Value.(*TreeNode).Left)
// 			}
// 			if node.Value.(*TreeNode).Right != nil{
// 				list.PushBack(node.Value.(*TreeNode).Right)
// 			}
// 		}
// 		out = append(out,currentLevel)
// 	}
// 	return out
// }

