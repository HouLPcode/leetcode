/*
 * @lc app=leetcode id=332 lang=golang
 *
 * [332] Reconstruct Itinerary
 */
 //测试集 [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
// 测试集 [["JFK","KUL"],["JFK","NRT"],["NRT","JFK"]] ，这个顺序不能只看孩子节点的字母顺序，先访问KUL会导致后续没有路径可走了
 // 题目中要求存储的是边，不是节点
func findItinerary(tickets [][]string) []string {
	// DFS
	rtn := []string{"JFK"}//注意，起点需要提前存储，dfs函数中只添加to内容
	visited := make(map[string]bool)
	// visited["JFK"] = true
	from := "JFK"
	dfs(tickets, from, visited, &rtn)
	return rtn
}

func dfs(tickets [][]string, from string, visited map[string]bool, rtn *[]string){//此处的rtn是否需要是指针
	tos := findTos(tickets, from, visited)
	for _,to := range tos{
		if _,ok := visited[from+to]; !ok{
			visited[from+to] = true 
			(*rtn) = append((*rtn), to)
			dfs(tickets, to, visited, rtn)
		}
	}
}

func findTos(tickets [][]string,from string, visited map[string]bool) []string{
	tos := []string{}
	for _,v := range tickets{
		if v[0] == from && visited[from+v[1]] == false{
			tos = append(tos, v[1])
		}
	}
	sort.Strings(tos) // 排序,******按照排序顺序，有可能不能遍历图中所有的节点
	return tos
}

func findTo(tickets [][]string,from string, visited map[string]bool) string{
	//找到from去的地方
	tos := []string{} // TODO 优化，过程中找到值
	for _,v := range tickets{
		if v[0] == from {
			if _,ok := visited[v[0]+v[1]]; !ok{
				//to取较小的
				tos = append(tos,v[1])
			}	
		}
	}
	if len(tos) == 0{
		return ""
	}
	to := tos[0]//肯定有值，不越界
	for _,v := range tos{
		if strings.Compare(to,v) == 1{ //顺序符合题目中的要求吗？？？
			to = v
		}
	}
	return to
}

