/*
 * @lc app=leetcode id=915 lang=golang
 *
 * [915] Partition Array into Disjoint Intervals
 */
// err --------------------------------
// [1,1]
// [90,47,69,10,43,92,31,73,61,97]
// [12,75,26,38,56,59,83,55,49,52,27,48,77,21,27,79,54,15,59,22,34,35,81,67,2,41,40,0,73,61,75,8,86,42,49,83,43,16,2,54,26,35,15,63,31,24,51,86,6,35,42,37,83,51,34,21,71,57,61,76,50,1,43,32,19,13,67,87,3,33,38,34,34,84,38,76,52,7,27,49,2,78,56,28,70,6,64,87,100,97,99,97,97,100,100,100,97,89,98,100]
// man
func partitionDisjoint(A []int) int {
	// -> max
	maxs := []int{} // 索引
	max := -1
	for i := 0; i < len(A); i++ {
		if A[i] > max {
			maxs = append(maxs, i)
			max = A[i]
		}
	}
	// <- min
	minMap := make(map[int]bool, 0) // 索引
	min := 1000001
	for i := len(A) - 1; i >= 0; i-- {
		if A[i] < min {
			min = A[i]
			minMap[i+1] = true
		}
	}
	rtn := 0
	for _, v := range maxs {
		if minMap[v] == true {
			rtn = v
			break
		}
	}
	if rtn == 0 {
		return 1
	}
	return rtn
}
