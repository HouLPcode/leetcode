/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */
func moveZeroes(nums []int)  {
	//两个指针
	p0,px := 0,0
	for {
		//找p0,p0前面都是处理完元素
		for ; p0 < len(nums) && nums[p0] != 0; p0++{}
		// P0和px中间全是0
		for px=p0+1; px < len(nums) && nums[px]==0;px++{}
		if p0 == len(nums) || px == len(nums){
			break
		}else{
			nums[p0],nums[px] = nums[px],nums[p0]
			p0++
			px++//此处的px会被找px的循环覆盖掉(px=p0+1)
		}
	}
}

