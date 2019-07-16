/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */
//  0 ms
// [2,0,1]
// O(n) 两个指针， 保证p0之前的数字全是0，p2后面的数字全是2，cur为扫描指针，p0和cur之间全是1，cur和p2之间是需要继续探索的数字
// 终止条件 cur > p2 //等于的时候表示除了当前访问的cur，没有剩余的未访问节点了
func sortColors(nums []int) {
	p0, p2 := 0, len(nums)-1
	// 长度不够3没影响？？？？
	for ; p0 < len(nums) && nums[p0] == 0; p0++ {
	} // 保证p0前面都是0，不包含p0
	for ; p2 >= 0 && nums[p2] == 2; p2-- {
	}
	for cur := p0; cur <= p2; { // 注意此处有等号
		if nums[cur] == 0 {
			nums[p0], nums[cur] = nums[cur], nums[p0] // cur <-> p0
			p0++
			cur++ // 注意此处移动cur
		} else if nums[cur] == 1 {
			cur++
		} else {
			nums[cur], nums[p2] = nums[p2], nums[cur]
			p2-- // 注意此处不知道交换前p2的值是什么，所以此处不需要
			// cur++ 一定注意，此处不需要
		}
	}
}
