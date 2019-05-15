/*
 * @lc app=leetcode id=372 lang=golang
 *
 * [372] Super Pow
 */
func superPow(a int, b []int) int {
	//不考虑越界问题，指定的是int
	// a==1 加速？？？
	// 思路: 把b组成int，然后折半相乘
	m := slice2int(b)
	return pow(a,m)%1337
}

func pow(x,m int)int{
	if m == 1{
		return x
	}else if m == 0{
		return 1
	}
	if m&1 == 0{
		return pow(x,m/2)*pow(x,m/2)
	}else{
		return pow(x,m/2)*pow(x,m/2)*x
	}
}

func slice2int(nums []int)int{
	cnt := 0
	for _,v := range nums{
		cnt = cnt * 10 + v
	}
	return cnt
}
