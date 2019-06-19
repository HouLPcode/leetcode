/*
 * @lc app=leetcode id=784 lang=golang
 *
 * [784] Letter Case Permutation
 */
//  76 ms 72.62 %
// 该题目中的字母顺序是固定好的，不能 0->len(S) 循环改变顺序
func letterCasePermutation(S string) []string {
	rtn := []string{}
	dfs(S, 0, []byte{}, &rtn)
	return rtn
}

// used[i][0] 代表小写 used[i][1]代表大写， 数字直接赋值两个数
func dfs(candidates string, i int, curr []byte, ans *[]string)  {
	if len(curr) == len(candidates) {
		(*ans) = append((*ans), string(curr))
		return
	}

	if candidates[i] >='0' && candidates[i] <= '9' {
		curr = append(curr, candidates[i])
		dfs(candidates, i+1, curr, ans)
	} else if candidates[i] >='a' && candidates[i] <= 'z'{
		curr = append(curr, candidates[i])
		dfs(candidates, i+1, curr, ans)
		curr = curr[:len(curr)-1]

		curr = append(curr, candidates[i]-'a'+'A')
		dfs(candidates, i+1, curr, ans)
		curr = curr[:len(curr)-1]
	}else if candidates[i] >='A' && candidates[i] <= 'Z'{
		curr = append(curr, candidates[i])
		dfs(candidates, i+1, curr, ans)
		curr = curr[:len(curr)-1]

		curr = append(curr, candidates[i]-'A'+'a')
		dfs(candidates, i+1, curr, ans)
		curr = curr[:len(curr)-1]
	}
}

