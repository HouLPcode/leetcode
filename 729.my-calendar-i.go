/*
 * @lc app=leetcode id=729 lang=golang
 *
 * [729] My Calendar I
 */
//  96 ms 65.96 %
//  暴力求解，通过[][2]int缓存book区间
// 比如两次订阅[s1,e1) [s2,e2)
// 		不冲突 e1<=s2 || e2<=s1
// 		冲突时 s2<e1 && s1<e2
type MyCalendar struct {
    Buf [][2]int
}


func Constructor() MyCalendar {
    return MyCalendar{
		make([][2]int, 0),
	}
}


func (this *MyCalendar) Book(start int, end int) bool {
    for _,v := range this.Buf {
		if start < v[1] && v[0] < end { // 冲突
			return false
		}
	}
	this.Buf = append(this.Buf, [2]int{start, end})
	return true
}


/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */

