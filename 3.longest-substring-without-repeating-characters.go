/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */
 // 两个指针组成子序列，map中缓存字母出现的位置，不是出现过没
// ✘ testcase: '"abba"'
func lengthOfLongestSubstring(s string) int {
	count := 0
	// alpMap := make(map[byte]bool, 0)
	alpMap := make(map[byte]int, 128)// byte最多128个，string长度不超过intMAX
	for p1,p2:=0,0; p2<len(s); p2++{//只循环右指针，左指针通过map缓存进行赋值
		if val,ok := alpMap[s[p2]]; ok{ // 字符出现过
			//如果此处应该把p1到val+1出现的字母从map中删除，则下面这句判断可以不写
			if val+1 > p1 { //***一定注意此处判断：重复字母的出现位置可能在p1之前，此时p1p2构成的子串中没有重复元素，p1不用移动
				p1 = val+1//左指针改为重复字符后一个字符，p1到val+1之间的字母虽然在map中仍然存在，但是在子串中是不存在的，无效的。p1只增不减	
		   }
		}
		alpMap[s[p2]] = p2//添加或更新字母出现的位置
		if p2-p1+1 > count{
			count = p2-p1+1
		}
	}
	return count
}
