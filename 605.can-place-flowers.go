/*
 * @lc app=leetcode id=605 lang=golang
 *
 * [605] Can Place Flowers
 */
//  (16 ms) √ Your runtime beats 88.04 %
// [0]\n1
// [0,0,1,0,0]\n1
// [0,0,0,0,0,1,0,0]\n0
func canPlaceFlowers(flowerbed []int, n int) bool {
	// 贪心思想，直接遍历
	if n == 0 {
		return true
	}
	i := 0
	for i+1 < len(flowerbed) {
		if i == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 { //0
			flowerbed[i] = 1
			n--
			if n == 0 {
				return true
			}
			i += 2
		} else if i > 0 && flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			n--
			if n == 0 {
				return true
			}
			i += 2
		} else {
			i++
		}
	}
	if i == len(flowerbed)-1 && flowerbed[i] == 0 {
		if i == 0 {
			n--
		} else if i-1 >= 0 && flowerbed[i-1] == 0 {
			n--
		}
	}
	if n == 0 {
		return true
	}
	return false
}
