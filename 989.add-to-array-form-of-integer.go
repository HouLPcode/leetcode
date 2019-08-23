/*
 * @lc app=leetcode id=989 lang=golang
 *
 * [989] Add to Array-Form of Integer
 */
// 220 ms, faster than 61.74%
// [6]\n809
func addToArrayForm(A []int, K int) []int {
	flag := 0 //进位标值
	i := len(A) - 1
	for ; i >= 0; i-- { // 逆序加法
		// if K == 0 { 不用退出，剩下的用来处理flag进位
		// 	break
		// }
		flag, A[i] = (A[i]+K%10+flag)/10, (A[i]+K%10+flag)%10
		K /= 10
	}
	tmp := []int{}
	var d int
	for K != 0 {
		flag, d = (K%10+flag)/10, (K%10+flag)%10 // 注意此处 d 不参与计算，d是为了保存当前值
		K /= 10
		tmp = append(tmp, d)
	}
	if flag != 0 {
		tmp = append(tmp, flag)
	}
	if len(tmp) == 0 {
		return A
	}

	rtn := make([]int, len(A)+len(tmp))
	for i := 0; i < len(tmp); i++ {
		rtn[i] = tmp[len(tmp)-1-i]
	}
	copy(rtn[len(tmp):], A)
	return rtn
}
