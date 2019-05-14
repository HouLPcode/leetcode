/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */
// [-1]\n2 
// [1,2]\n2 
// [1,2]\n3 
// 注意k比nums长度还大时候的翻转
func rotate(nums []int, k int)  {
	k %= len(nums)
	// 前部分翻转，后半部分翻转，整体翻转
	reverse(&nums,0,len(nums)-k-1)
	reverse(&nums,len(nums)-k,len(nums)-1)
	reverse(&nums,0,len(nums)-1)
}

func reverse(nums *[]int, start,end int){
	for ; start < end; start,end = start+1,end-1{
		(*nums)[start],(*nums)[end] = (*nums)[end],(*nums)[start]
	}
}

