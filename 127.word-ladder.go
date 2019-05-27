/*
 * @lc app=leetcode id=127 lang=golang
 *
 * [127] Word Ladder
 */
//  双向BFS
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// 特殊处理下endWord不在wordList中的情景
	if hasEndWord(endWord, wordList) == false{
		return 0
	}
	depth := 0
	visited := make(map[string]bool, len(wordList))

	// 使用两个map，表示头尾的遍历，不需要queue
	map1 := make(map[string]bool, 0)
	map2 := make(map[string]bool, 0)
	map1[beginWord] = true //首尾节点入栈
	map2[endWord] = true
	visited[beginWord] = true
	visited[endWord] = true

	for len(map1) != 0{
		depth++
		tmpMap := make(map[string]bool, 0)// 存储子节点
		for k := range map1{
			nexts := connect(k, wordList)
			for _,next := range nexts{
				if map2[next] == true{ // 首尾相遇，返回深度
					return depth+1
				}
				if visited[next] == false{
					tmpMap[next] = true
					visited[next] = true
				}
			}
			delete(map1, k)
		}
		map1, map2 = map2, tmpMap // map怎么交换？？？？？？
	}
	return 0
}

func hasEndWord(s string, wordList []string) bool {
	for _,v := range wordList{
		if s == v{
			return true
		}
	}
	return false
}

func connect(word string, wordList []string) []string {
	// 返回和word只有1字符之差的所有字符串
	result := make([]string, 0, len(wordList))
	for _, str := range wordList {
		if discOneChar(word, str) {
			result = append(result, str)
		}
	}
	return result
}

func discOneChar(s, t string) bool {
	// s,t 长度相同，是否只有1个字符不同
	cnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			cnt++
			if cnt > 1{
				return false
			}
		}
	}
	return cnt == 1
}
