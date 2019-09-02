/*
 * @lc app=leetcode.cn id=735 lang=golang
 *
 * [735] 行星碰撞
 */
// 8 ms, faster than 100.00%
// [-2,-2,1,-2]
// [8, -8]
// 法二： 最终结果一定是 -----+++++++,其中负数在左，可以为空
// stack 模拟过程
// +， 直接push
// -
// 1. len == 0 ,push
// 2. top == -, push
// 3. top == + ， 循环 top+v > 0 , == 0 也退出
func asteroidCollision(asteroids []int) []int {
	// 只有负数才有可能碰撞
	stack := make([]int, 0, len(asteroids))

	for i := 0; i < len(asteroids); i++ {
		v := asteroids[i]
		if v > 0 {
			stack = append(stack, v)
		} else { // v == -
			if len(stack) == 0 {
				stack = append(stack, v)
			} else {
				top := stack[len(stack)-1]
				if top < 0 {
					stack = append(stack, v)
				} else { // top > 0
					if top+v == 0 {
						stack = stack[:len(stack)-1] //pop
					} else if top+v < 0 {
						stack = stack[:len(stack)-1] //pop
						i--                          // 表示v依然有效，重新入队，避免循环
					} else {
						//...
					}
				}
			}
		}
	}
	return stack
}
