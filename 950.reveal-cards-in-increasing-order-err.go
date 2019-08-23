import "sort"

/*
 * @lc app=leetcode id=950 lang=golang
 *
 * [950] Reveal Cards In Increasing Order
 */
// err-----------------------------------------
// [1,2,3,4,5,6]
// [1,2,3,4,5,6,7,8,9]
func deckRevealedIncreasing(deck []int) []int {
	// 先整体排序
	// 前半部分间隔一个排在一起， 0，2，4，6，8。。。 n = (l+1)/2
	// 后半部分右旋转 n次， n是左半部分的长度
	sort.Ints(deck)
	right := make([]int, len(deck)/2) // 右半部分
	copy(right, deck[(len(deck)+1)/2:])
	//间隔一个填充左半部分
	j := (len(deck) - 1) / 2
	i := 0
	if len(deck)&1 == 0 {
		i = len(deck) - 2
	} else {
		i = len(deck) - 1
	}
	for ; i >= 0; i -= 2 {
		deck[i] = deck[j]
		j--
	}
	if len(deck)&1 == 0 {
		for i := 1; i < len(deck); i += 2 {
			deck[i] = right[0]
			right = right[1:]
			if len(right) == 0 {
				break
			}
			right = append(right, right[0])
			right = right[1:]
		}
	} else {
		for i := 1; i < len(deck); i += 2 {
			// 有旋转一次，然后弹出 [0]元素
			right = append(right, right[0])
			deck[i] = right[1]
			right = right[2:]
		}
	}

	return deck
}
