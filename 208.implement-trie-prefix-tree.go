/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 */
 // 64 ms, faster than 83.33%
type Trie struct {
	Alphs [26]*Trie
	IsWord bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{
		Alphs:[26]*Trie{},
		IsWord:false,
	}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	curNode := this
	// range string 
	for _,v := range []byte(word){
		// Trie 嵌套怎么实现
		if curNode.Alphs[v-'a'] == nil { // 没有这个字母
			curNode.Alphs[v-'a'] = &Trie{Alphs:[26]*Trie{},}
		}
		curNode = curNode.Alphs[v-'a']
	}
	// 赋值完之后，怎么对最后一个Trie.IsWord赋值true
	curNode.IsWord = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	curNode := this
	for _, v := range []byte(word){
		if curNode.Alphs[v-'a'] == nil {
			return false
		}
		curNode = curNode.Alphs[v-'a']
	}
	return curNode.IsWord
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    curNode := this
	for _, v := range []byte(prefix) {
		if curNode.Alphs[v-'a'] == nil {
			return false
		}
		curNode = curNode.Alphs[v-'a']
	}
	return true
}

// TODO 怎么输出trie中的所有单词？

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

