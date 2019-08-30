/*
 * @lc app=leetcode id=919 lang=golang
 *
 * [919] Complete Binary Tree Inserter
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // 32 ms, faster than 63.16% 
// 保存父节点列表进行快速插入
type CBTInserter struct {
	Root *TreeNode
	Parents []*TreeNode
}


func Constructor(root *TreeNode) CBTInserter {
	// root 是一棵树，并不是只有1个节点
	rtn := CBTInserter{}
	rtn.Root = root
	// 插入parents节点
	rtn.Parents = append(rtn.Parents, root)
	for len(rtn.Parents) > 0 {
		size := len(rtn.Parents)
		i:=0
		for ; i<size; i++ {
			if rtn.Parents[i].Left == nil ||
				rtn.Parents[i].Right == nil {
				break
			}
			rtn.Parents = append(rtn.Parents,rtn.Parents[i].Left,rtn.Parents[i].Right)
		}
		if i < size {
			rtn.Parents = rtn.Parents[:size]	// 找到倒数第二行
			break
		} else {
			rtn.Parents = rtn.Parents[size:] // 新的一行
		}
	}

	return rtn
}


func (this *CBTInserter) Insert(v int) int {
	i := 0
	for ; i < len(this.Parents); i++{
		if this.Parents[i].Left != nil &&
			this.Parents[i].Right != nil{
			continue
		} else if this.Parents[i].Left == nil{
			this.Parents[i].Left  = &TreeNode{v,nil,nil,}
			return this.Parents[i].Val
		} else {
			this.Parents[i].Right  = &TreeNode{v,nil,nil,}
			return this.Parents[i].Val
		}
	}
	if i == len(this.Parents) {
		// 新的一行
		size := len(this.Parents)
		for i:=0;i< size; i++ {
			// 注意此处是0
			this.Parents = append(this.Parents,this.Parents[0].Left,this.Parents[0].Right)
			this.Parents = this.Parents[1:]
		}
		this.Parents[0].Left = &TreeNode{v,nil,nil,}
		return this.Parents[0].Val
	}
	return 0
}


func (this *CBTInserter) Get_root() *TreeNode {
	return this.Root
}


/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(v);
 * param_2 := obj.Get_root();
 */
