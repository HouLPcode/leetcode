/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */
func groupAnagrams(strs []string) [][]string {
	//法1 hash表  sortstr -> []string 对string排序？？？
	//法2 hash表  "[0 1 22]" ->[]string 下面是该方法
	result := [][]string{}
	numMap := make(map[string][]string, 0)
	for _,str := range strs{
		alpnum := make([]int, 26)//26个小写字母
		for i:=0; i<len(str); i++{
			alpnum[int(str[i]-'a')]++
		}
		alpKey := fmt.Sprint(alpnum)
		if _,ok := numMap[alpKey]; ok{
			numMap[alpKey] = append(numMap[alpKey], str)
		}else{
			numMap[alpKey] = []string{str}
		}
	}
	for _,v := range numMap{
		result = append(result, v)
	}

	return result
}

