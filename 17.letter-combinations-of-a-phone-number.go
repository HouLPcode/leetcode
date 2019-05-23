/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */
func letterCombinations(digits string) []string {
	if len(digits) == 0{
		return []string{}
	}
	result := make([]string, 0)
	generate(0, digits, "", &result)
	return result
}
var alhMap = map[byte][]string{
	'2':[]string{"a","b","c"},
	'3':[]string{"d","e","f"},
	'4':[]string{"g","h","i"},
	'5':[]string{"j","k","l"},
	'6':[]string{"m","n","o"},
	'7':[]string{"p","q","r","s"},
	'8':[]string{"t","u","v"},
	'9':[]string{"w","x","y","z"},
}

func generate(k int, digits string,strTemp string, result *[]string){
	if k == len(digits){
		*result = append((*result), strTemp)
		return
	}
	strTempCopy := strTemp //前备份
	for _,v := range alhMap[digits[k]]{
		generate(k+1, digits, strTemp+v, result)
		strTemp = strTempCopy// 后恢复
	}
}

