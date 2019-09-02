/*
 * @lc app=leetcode.cn id=735 lang=golang
 *
 * [735] 行星碰撞
 */
// 12 ms, faster than 84.00%
// [-2,-2,1,-2]
// [8, -8]
// stack 模拟
func asteroidCollision(asteroids []int) []int {
	// 只有负数才有可能碰撞
	stack := make([]int, 0, len(asteroids))

	for _, v := range asteroids {
		if v > 0 || len(stack) == 0 || stack[len(stack)-1] < 0 {
			stack = append(stack, v)
		} else {
			for {
				if len(stack) == 0 {
					stack = append(stack, v)
					break
				} else {
					top := stack[len(stack)-1]
					if top < 0 {
						stack = append(stack, v)
						break // 同方向，退出
					} else if top+v == 0 {
						stack = stack[:len(stack)-1] // pop
						break
					} else if top+v > 0 {
						break
					} else if top+v < 0 {
						stack = stack[:len(stack)-1] //pop
					}
				}
			}
		}
	}
	return stack
}
