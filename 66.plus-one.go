/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */
 //注意输入为空的情况，999的情况
func plusOne(digits []int) []int {
    if len(digits) == 0{
        return []int{1}
    }
    // if digits[len(digits)-1] < 9 { //加速？？？
    //     digits[len(digits)-1]++
    //     return digits
    // }
    //有进位
    var cf int
    digits[len(digits)-1], cf = (digits[len(digits)-1] + 1) % 10,(digits[len(digits)-1] + 1) / 10
    for i:= len(digits)-2; i>=0;i--{
        // 注意此处需要同时赋值
        digits[i], cf = (digits[i] + cf) % 10,(digits[i] + cf) / 10
    }
    if cf == 1{
        digits = append([]int{1},digits...)
    }
    return digits
}

