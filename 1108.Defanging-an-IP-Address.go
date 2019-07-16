// 0 ms
func defangIPaddr(address string) string {
	rtn := []byte(address)
	rtn = append(rtn, []byte{'1', '2', '3', '4', '5', '6'}...)
	for i, j := len(rtn)-7, len(rtn)-1; i >= 0; i-- {
		if rtn[i] == '.' {
			rtn[j] = ']'
			j--
			rtn[j] = '.'
			j--
			rtn[j] = '['
			j--
		} else {
			rtn[j] = rtn[i]
			j--
		}
	}
	return string(rtn)
}