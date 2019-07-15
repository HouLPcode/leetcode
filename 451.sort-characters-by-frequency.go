import "sort"

/*
 * @lc app=leetcode id=451 lang=golang
 *
 * [451] Sort Characters By Frequency
 */
//  4 ms 95.60%
// 通过128字节数组缓存数据 byte -> f
// 按照f排序
func frequencySort(s string) string {
	charfre := make([]int, 128)
	for _, v := range []byte(s) {
		charfre[int(v)]++
	}
	sorts := []Data{}

	for k, v := range charfre {
		if v > 0 {
			sorts = append(sorts, Data{byte(k), v})
		}
	}
	sort.Sort(SortMap(sorts))

	rtn := []byte{}
	for _, v := range sorts {
		for i := v.Fre; i > 0; i-- {
			rtn = append(rtn, v.Char)
		}
	}
	return string(rtn)
}

type Data struct {
	Char byte
	Fre  int
}

type SortMap []Data

func (s SortMap) Len() int { return len(s) }

func (s SortMap) Less(i, j int) bool { // 注意从多到少
	return s[i].Fre > s[j].Fre
}

func (s SortMap) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
