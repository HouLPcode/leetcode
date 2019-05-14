/*
 * @lc app=leetcode id=504 lang=golang
 *
 * [504] Base 7
 */
func convertToBase7(num int) string {
	mflag := true //正数
	if num < 0{
		mflag = false//负数
		num = -num // [-1e7, 1e7],不越界吧？？？
	} 
	
	str := ""
	str = fmt.Sprint(num%7) + str // 注意num=0的处理
	num /= 7
	for num != 0{
		str = fmt.Sprint(num%7) + str
		num /= 7
	}
	if !mflag{
		str = "-" + str
	}
	return str
}

