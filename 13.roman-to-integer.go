/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */
package main

// 方法1.遍历的时候，每次尽可能匹配两个字符，不匹配的时候再当1个字符处理
// 方法2. 技巧s[i]<s[i+1] ,则sum-s[i]+s[i+1]
// 注意最后一个字符的处理
// 注意只有一个字符的处理
func romanToInt(s string) int {
	nummap := map[byte]int{//缓存map
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	cnt := 0
	for i:=0;i<len(s)-1;i++{
		if nummap[s[i]] < nummap[s[i+1]]{//i+2取得是开区间，不越界吧？？？
			cnt -= nummap[s[i]]
		}else{
			cnt += nummap[s[i]]
		}
	}
	cnt += nummap[s[len(s)-1]]
	return cnt
}
