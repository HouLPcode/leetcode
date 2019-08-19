import "strconv"

/*
 * @lc app=leetcode id=537 lang=golang
 *
 * [537] Complex Number Multiplication
 */
// "78+-76i" "-86+72i"
func complexNumberMultiply(a string, b string) string {
	nums1 := conv2int(a)
	nums2 := conv2int(b)
	z := nums1[0]*nums2[0] - nums1[1]*nums2[1] // 整
	f := nums1[0]*nums2[1] + nums1[1]*nums2[0] // 辅
	return strconv.Itoa(z) + "+" + strconv.Itoa(f) + "i"
}

func conv2int(a string) (rtn [2]int) {
	i := 0
	num := int(0)
	if a[i] == '-' { // 不要忘了第一个数的正负
		i++
		for ; i < len(a) && a[i] != '+'; i++ {
			num *= 10
			num += int(a[i] - '0')
		}
		rtn[0] = -num
	} else {
		for ; i < len(a) && a[i] != '+'; i++ {
			num *= 10
			num += int(a[i] - '0')
		}
		rtn[0] = num
	}

	i++ // +

	if a[i] == '-' {
		num = 0
		i++
		for ; i < len(a) && a[i] != 'i'; i++ {
			num *= 10
			num += int(a[i] - '0')
		}
		num = -num
	} else {
		num = 0
		for ; i < len(a) && a[i] != 'i'; i++ {
			num *= 10
			num += int(a[i] - '0')
		}
	}
	rtn[1] = num
	return rtn
}
