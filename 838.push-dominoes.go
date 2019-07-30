/*
 * @lc app=leetcode id=838 lang=golang
 *
 * [838] Push Dominoes
 */
//  4 ms) √ Your runtime beats 100 %
func pushDominoes(dominoes string) string {
	doms := []byte(dominoes)
	p0 := 0
	for p0 < len(doms) {
		if doms[p0] == '.' {
			p0++
		} else if doms[p0] == 'L' {
			for i := p0 - 1; i >= 0 && doms[i] == '.'; i-- {
				// 左边所有连续 . -> L
				doms[i] = 'L'
			}
			p0++ // 注意p0右移
		} else { // p0 == R
			p1 := p0 + 1
			for ; p1 < len(doms) && doms[p1] == '.'; p1++ {
			}
			if p1 == len(doms) {
				// R
				for i := p0 + 1; i < len(doms); i++ {
					doms[i] = 'R'
				}
				break
			} else if doms[p1] == 'L' {
				// ... = RL
				for i, j := p0, p1; i < j; i, j = i+1, j-1 {
					doms[i] = 'R'
					doms[j] = 'L'
				}
				p0 = p1 + 1
			} else { // R
				// p0 ... p1 = R
				for i := p0 + 1; i < p1; i++ {
					doms[i] = 'R'
				}
				p0 = p1 //  此处p0不右移
			}
		}
	}
	return string(doms)
}
