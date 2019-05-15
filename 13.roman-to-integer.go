/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */
package main

// 遍历的时候，每次尽可能匹配两个字符，不匹配的时候再当1个字符处理
// 注意最后一个字符的处理
// 注意只有一个字符的处理
func romanToInt(s string) int {
	nummap := map[string]int{//缓存map
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
		"IV":4,
		"IX":9,
		"XL":40,
		"XC":90,
		"CD":400,
		"CM":900,
	}
	cnt := 0
	for i:=0;i<len(s)-1;i++{//少循环一次，避免i+1越界
		if v,ok := nummap[s[i:i+2]]; ok{//i+2取得是开区间，不越界吧？？？
			cnt += v
			i++
		}else{
			cnt += nummap[s[i:i+1]]
		}
	}
	// 处理最后1个字符
	if len(s) == 1{
		return nummap[s]
	}
	if _,ok := nummap[s[len(s)-2:]]; !ok{ // -2可能越界
		cnt += nummap[s[len(s)-1:]]
	}
	return cnt
}
