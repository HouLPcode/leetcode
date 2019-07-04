/*
 * @lc app=leetcode id=500 lang=golang
 *
 * [500] Keyboard Row
 */
//  0 ms
func findWords(words []string) []string {
	row := make(map[byte]int, 52)
	row['Q'] = 1
	row['W'] = 1
	row['E'] = 1
	row['R'] = 1
	row['T'] = 1
	row['Y'] = 1
	row['U'] = 1
	row['I'] = 1
	row['O'] = 1
	row['P'] = 1
	row['A'] = 2
	row['S'] = 2
	row['D'] = 2
	row['F'] = 2
	row['G'] = 2
	row['H'] = 2
	row['J'] = 2
	row['K'] = 2
	row['L'] = 2
	row['Z'] = 3
	row['X'] = 3
	row['C'] = 3
	row['V'] = 3
	row['B'] = 3
	row['N'] = 3
	row['M'] = 3
	row['q'] = 1
	row['w'] = 1
	row['e'] = 1
	row['r'] = 1
	row['t'] = 1
	row['y'] = 1
	row['u'] = 1
	row['i'] = 1
	row['o'] = 1
	row['p'] = 1
	row['a'] = 2
	row['s'] = 2
	row['d'] = 2
	row['f'] = 2
	row['g'] = 2
	row['h'] = 2
	row['j'] = 2
	row['k'] = 2
	row['l'] = 2
	row['z'] = 3
	row['x'] = 3
	row['c'] = 3
	row['v'] = 3
	row['b'] = 3
	row['n'] = 3
	row['m'] = 3

	rtn := make([]string, 0, len(words))
	for _, word := range words {
		for alph, i := row[word[0]], 0; i < len(word); i++ {
			if row[word[i]] != alph {
				break
			} else if i == len(word)-1 {
				rtn = append(rtn, word)
			}
		}
	}
	return rtn
}

