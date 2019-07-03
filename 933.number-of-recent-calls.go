/*
 * @lc app=leetcode id=933 lang=golang
 *
 * [933] Number of Recent Calls
 */
//  148 ms 59.82 %
type RecentCounter struct {
	buf []int // 缓存t
}

func Constructor() RecentCounter {
	return RecentCounter{
		buf: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.buf = append(this.buf, t)

	for t-3000 > this.buf[0] { // 不会越界
		this.buf = this.buf[1:]
	}
	return len(this.buf)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

