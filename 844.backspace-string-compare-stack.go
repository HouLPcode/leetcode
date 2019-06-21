/*
 * @lc app=leetcode id=844 lang=golang
 *
 * [844] Backspace String Compare
 */
//  0 ms
//  stack处理，待优化
// "y#fo##f"\n"y#f#o##f"
func backspaceCompare(S string, T string) bool {
	stack1 := make([]byte, 0, len(S))
	for i:=0; i< len(S); i++ {
		if S[i] == '#' { // 一定注意此处分析问题要全面
			if len(stack1) > 0 { //注意此处还有S[i]==#,len()==0的情况
				stack1 = stack1[:len(stack1)-1]
			}
		} else {
			stack1 = append(stack1, S[i])
		}
	}
	stack2 := make([]byte, 0, len(T))
	for i:=0; i < len(T); i++ {
		if T[i] == '#' {
			if len(stack2) > 0 {
				stack2 = stack2[:len(stack2)-1]
			}
		}else {
			stack2 = append(stack2, T[i])
		}
	}
	return reflect.DeepEqual(stack1, stack2)
}

