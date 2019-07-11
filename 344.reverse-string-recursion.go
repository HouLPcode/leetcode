/*
 * @lc app=leetcode id=344 lang=golang
 *
 * [344] Reverse String
 */
 func reverseString(s []byte)  {// slice是值引用？？？
    // recursion 递归实现
    if len(s) < 2 {
        return 
    }
    s[0], s[len(s)-1] = s[len(s)-1], s[0]
    reverseString(s[1:len(s)-1])
}


