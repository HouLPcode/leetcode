/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

func isAnagram(s string, t string) bool {
	m1 := make(map[rune]int,26)//一共有26个小写字母
	m2 := make(map[rune]int, 26)
	// range string 返回的是 rune
	for _,v := range s{
		m1[v] += 1 // ++  ???
	}
	for _,v := range t{
		m2[v] += 1
	} 
	return reflect.DeepEqual(m1,m2)
	// return cmp.Equal(m1, m2) // import "github.com/google/go-cmp/cmp"

	// map == nil可以，其他的情况不支持直接 == 
	// map怎么比较是否相同？  
	// 进而组合结构如何比较是否相同，slice，string，interface？？？
}

