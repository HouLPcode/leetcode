/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */
func canJump(nums []int) bool {
    // 贪心算法尝试
    // 技巧：从右往左遍历，每次记录能够到达的终点的最左边的值lastPosition，初始时为最后一个索引
    lastPosition := len(nums)-1
    for i:=len(nums)-1; i>=0; i-- {
        if i+nums[i] >= lastPosition{//能够到达最后，更新最左边的位置
            lastPosition = i
        }
    }
    return lastPosition == 0
}

