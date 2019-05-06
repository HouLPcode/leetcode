/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 */
type Trie struct {
	Alphs map[rune]*Trie//??????
	IsWord bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{
		Alphs:make(map[rune]*Trie),
		IsWord:false,//注意每个结尾都有逗号，否则编译报语法错误
	}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	curNode := this
	// range string 
	for _,v := range word{
		// Trie 嵌套怎么实现
		// 注意if后面用的是 ;
		if _,ok := curNode.Alphs[v]; !ok{
			// 这一句必须有，此时的map还没有make？？？Constructor()函数不是make了吗？？？？
			// curNode.Alphs = make(map[rune]Trie)
			newNode := Constructor()
			curNode.Alphs[v] = &newNode
			curNode = &newNode
		}
	}
	// 赋值完之后，怎么对最后一个Trie.IsWord赋值true
	curNode.IsWord = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	curNode := this
	for _, v := range word{
		if _, ok := curNode.Alphs[v]; !ok{
			return false
		}
		curNode = curNode.Alphs[v]
	}
	return curNode.IsWord
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    curNode := this
	for _, v := range prefix{
		if _, ok := curNode.Alphs[v]; !ok{
			return false
		}
		curNode = curNode.Alphs[v]
	}
	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

