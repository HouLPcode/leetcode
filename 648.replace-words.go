/*
 * @lc app=leetcode id=648 lang=golang
 *
 * [648] Replace Words
 */
 // 20 ms, faster than 90.00% 
 // 暴力解决
 // dict 排序后， 对句子中的每个单词判断前缀
func replaceWords(dict []string, sentence string) string {
	sort.Strings(dict)
	words := strings.Split(sentence, " ")
	for i:=0; i<len(words); i++ {
		for j:=0; j<len(dict); j++{
			if strings.HasPrefix(words[i], dict[j]){
				words[i] = dict[j]
				break
			}
		}
	}
	return strings.Join(words, " ")
}

