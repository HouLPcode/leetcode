import (
	"sort"
	"strings"
)

/*
 * @lc app=leetcode id=937 lang=golang
 *
 * [937] Reorder Log Files
 */
//  4 ms, faster than 93.83%
// ["w 7 2", "l 1 0", "6 066", "o aay", "e yal", "b aay"] 字母排序错误
func reorderLogFiles(logs []string) []string {
	// 遍历一遍，把数字日志排序到右侧，顺序不变
	pri2 := len(logs) - 1                 // pri2右侧的元素全是数字日志，不包含pri2
	for i := len(logs) - 1; i >= 0; i-- { // 不管是否交换，i都需要--
		if logs[i][len(logs[i])-1] >= '0' && logs[i][len(logs[i])-1] <= '9' {
			// 最后一个字符数字则时数字日志
			logs[i], logs[pri2] = logs[pri2], logs[i]
			pri2--
		}
	}
	// [0,pri2+1] 是字符日志

	sort.Slice(logs[:pri2+1], func(i, j int) bool {
		// 先按照内容字母顺序排序，再按照标识符排序
		index1 := find(logs[i]) // 第一个空格后的字符
		index2 := find(logs[j])
		b := strings.Compare(logs[i][index1:], logs[j][index2:])
		b1 := strings.Compare(logs[i][:index1], logs[j][:index2])
		return b == -1 || (b == 0 && b1 == -1)
	})
	return logs
}

// 返回第一个空格后面的index
func find(s string) int {
	i := 0
	for s[i] != ' ' {
		i++
	}
	return i + 1
}
