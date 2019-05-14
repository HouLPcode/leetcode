/*
 * @lc app=leetcode id=520 lang=golang
 *
 * [520] Detect Capital
 */
 //TODO 4ms 时间待优化
func detectCapitalUse(word string) bool {
	//输入非空，且只包含大小写字母
	//"mL"  "Ml" "us" "UK"
	mflag := isCapital(word[0])
	for k,b := range []byte(word){
		if mflag != isCapital(b) {
			if mflag == true && k == 1{ //此处必须满足第一位是大写，第二位小写才能更新
				mflag = isCapital(b)
			}else{
				return false
			}
		}
	}
	return true
}

func isCapital(s byte)bool{
	if s >= byte('A') && s <= byte('Z'){
		return true
	}
	return false
}

