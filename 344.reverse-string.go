/*
 * @lc app=leetcode id=344 lang=golang
 *
 * [344] Reverse String
 */
func reverseString(s []byte)  {
	//只能在原数组上面改，O(1)空间
	for i, j := 0, len(s)-1; i < j; i,j = i+1,j-1{
		s[i], s[j] = s[j], s[i]
	}
}

