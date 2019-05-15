/*
 * @lc app=leetcode id=12 lang=golang
 *
 * [12] Integer to Roman
 */
func intToRoman(num int) string {
	// 	"I": 1,
	// 	"IV":4,
	// 	"V": 5,
	// 	"IX":9,
	// 	"X": 10,
	// 	"XL":40,
	// 	"L": 50,
	// 	"XC":90,
	// 	"C": 100,
	// 	"CD":400,
	// 	"D": 500,
	// 	"CM":900,
	// 	"M": 1000,
	// num 循环减去可能的最大值，-1000,-900,-500...
	str := ""
	for num >=1000{
		str += "M"
		num -= 1000
	}
	for num >=900{
		str += "CM"
		num -= 900
	}
	for num >=500{
		str += "D"
		num -= 500
	}
	for num >=400{
		str += "CD"
		num -= 400
	}
	for num >=100{
		str += "C"
		num -= 100
	}
	for num >=90{
		str += "XC"
		num -= 90
	}
	for num >=50{
		str += "L"
		num -= 50
	}
	for num >=40{
		str += "XL"
		num -= 40
	}
	for num >=10{
		str += "X"
		num -= 10
	}
	for num >=9{
		str += "IX"
		num -= 9
	}
	for num >=5{
		str += "V"
		num -= 5
	}
	for num >=4{
		str += "IV"
		num -= 4
	}
	for num >=1{
		str += "I"
		num -= 1
	}
	return str
}

