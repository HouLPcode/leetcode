import "sort"

// 4 ms 0%
func relativeSortArray(arr1 []int, arr2 []int) []int {
	// map 缓存arr2， 计数
	arr2Map := make(map[int]int, len(arr2))
	for _, v := range arr2 {
		arr2Map[v] = 1 // 多计数一次
	}

	arr3 := make([]int, 0)
	for i := 0; i < len(arr1); i++ {
		if arr2Map[arr1[i]] > 0 {
			arr2Map[arr1[i]]++
		} else {
			arr3 = append(arr3, arr1[i])
		}
	}

	arr1 = arr1[:0]
	for _, v := range arr2 {
		for i := 0; i+1 < arr2Map[v]; i++ {
			arr1 = append(arr1, v)
		}
	}
	sort.Ints(arr3)
	arr1 = append(arr1, arr3...)
	return arr1
}
