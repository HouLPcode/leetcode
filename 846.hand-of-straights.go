/*
 * @lc app=leetcode id=846 lang=golang
 *
 * [846] Hand of Straights
 */
//(32 ms) √ Your runtime beats 100 %
// [1]\n1
// [1,1,2,2,3,3]\n3
// map int -> time  然后遍历int
func isNStraightHand(hand []int, W int) bool {
	if len(hand)%W != 0 {
		return false
	}

	treeMap := make(map[int]int)
	min, max := 1000000001, -1

	for i := 0; i < len(hand); i++ {
		treeMap[hand[i]]++
		if min > hand[i] {
			min = hand[i]
		}
		if max < hand[i] {
			max = hand[i]
		}
	}
	// hand[i]范围 [min,max]
	for i := min; i <= max; { // 扫描可能的每个值,注意，此处不用i++
		if treeMap[i] == 0 {
			i++ // 此处进行i的递增，因为 i可能有多个，需要遍历完后再进行下一个
			continue
		}
		for j := i; j < i+W; j++ { // 连续的W个数字次数--
			if treeMap[j] <= 0 {
				return false
			}
			treeMap[j]--
			if treeMap[j] == 0 {
				delete(treeMap, j)
			}
		}
	}
	return len(treeMap) == 0
}
