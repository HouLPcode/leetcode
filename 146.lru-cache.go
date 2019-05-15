/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */
type LRUCache struct {
	buf map[int]int
	stack []int
	capacity int
}
// 注意 get操作会提高相关key的位置
// 注意 put操作可能是已经存在的key

// 数据k-v格式通过map[int][int]实现
// 顺序通过[]int，里面存储的是key
// get直接操作map, get操作会提高相关key的位置
// put 超过大小后，去除栈低元素，map删除相关key
func Constructor(capacity int) LRUCache {
    return LRUCache{
		buf:make(map[int]int), // map 大小？？？？？？
		stack:make([]int, 0, capacity),
		capacity:capacity,
	}
}


func (this *LRUCache) Get(key int) int {
    if rnt,ok := this.buf[key]; ok{
		//提高key在stack中的位置
		for k,v := range this.stack{
			if v == key{
				if k != len(this.stack)-1 {//栈顶不用更新
					// this.stack = this.stack[:k]+this.stack[k+1:]//k+1越界,且slice不支持直接加法
					this.stack = append(this.stack[:k],this.stack[k+1:]...)
					this.stack = append(this.stack,key)
				}
				break
			}
		}
		return rnt
	}
	return -1
}

//注意，put可以有重复key，此时只更新stack和map即可
func (this *LRUCache) Put(key int, value int)  {
	// 注意只能通过this调用Get方法
	if this.Get(key) != -1 { //已经存在key,stack中key已经在栈顶
		// this.buf[key] = value		
	}else if len(this.stack) == this.capacity{//没有key,且栈满
		delete(this.buf,this.stack[0])
		this.stack = this.stack[1:]//???没有移位，导致后面添加元素越界
		// this.stack[len(this.stack)] = key //append更耗时？？？ 该语句越界错误
		this.stack = append(this.stack, key)
		// this.buf[key] = value
	}else{//没有key，栈不满
		this.stack = append(this.stack, key)
		// this.buf[key] = value
	}
	this.buf[key] = value		
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

