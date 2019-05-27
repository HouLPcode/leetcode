/*
 * @lc app=leetcode id=127 lang=golang
 *
 * [127] Word Ladder
 */
//  由于需要找的是最短路径，所以要用BFS，第一次找到目标节点即为最短路径。DFS找到的不是最短
 func ladderLength(beginWord string, endWord string, wordList []string) int {
	depth := 0 //bfs怎么统计图的深度
	visited := make(map[string]bool, len(wordList))
	queue := list.New()

	//visited[beginWord] = true
	queue.PushBack(beginWord)
	for queue.Front() != nil {
		depth++
		for size := queue.Len(); size > 0; size-- { // *****注意此处的处理方式，一次处理一层
			node := queue.Front()
			queue.Remove(node)
			visited[node.Value.(string)] = true //出队列的时候标记为已访问
			if node.Value.(string) == endWord {
				return depth
			}
			for _, nextNode := range connect(node.Value.(string), wordList) {
				if visited[nextNode] == false {
					queue.PushBack(nextNode)
				}
			}
		}
	}
	return 0
}

func connect(word string, wordList []string) []string {
	// 返回和word只有1字符之差的所有字符串
	result := []string{}
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
		}
	}
	return cnt == 1
}