/*
 * @lc app=leetcode id=949 lang=golang
 *
 * [949] Largest Time for Given Digits
 */
//  0 ms
//  [0,0,0,0]
//  暴力破解， A(4,4) = 24种，由于肯定只有4个元素，所以不需要考虑DFS的排列，直接循环就能解决
func largestTimeFromDigits(A []int) string {
	maxMinutes := -1 //当前分钟数,一定注意此处不能初始化为0，否则无法判断返回"" 还是 "00:00"
	for i := 0; i < 4; i++ {
		if A[i] > 2 {
			continue
		}
		for j := 0; j < 4; j++ {
			if j == i || A[i]*10+A[j] > 23 {
				continue
			}
			for k := 0; k < 4; k++ {
				if k != i && k != j && A[k] < 6 && A[k]*10+A[6-i-j-k] < 59 {
					if (A[i]*10+A[j])*60+A[k]*10+A[6-i-j-k] > maxMinutes {
						maxMinutes = (A[i]*10+A[j])*60 + A[k]*10 + A[6-i-j-k]
					}
				}
			}
		}
	}
	if maxMinutes == -1 {
		return ""
	}
	return string([]byte{
		byte(maxMinutes/60/10) + '0',
		byte(maxMinutes/60%10) + '0',
		':',
		byte(maxMinutes%60/10) + '0',
		byte(maxMinutes%60%10) + '0',
	})
}
