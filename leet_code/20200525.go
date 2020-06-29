package main

import (
	"fmt"
	"time"
)

type LRUCache struct {
	Cap     int
	Data    []DataInfo
	DataMap map[int]int
}

func (l *LRUCache) Less(i int, j int) bool {
	return l.Data[i].UseTime > l.Data[j].UseTime
}

func (l *LRUCache) Len() int {
	return len(l.Data)
}

func (l *LRUCache) Swap(i int, j int) {
	l.Data[i], l.Data[j] = l.Data[j], l.Data[i]
}

func (l *LRUCache) Append(dataInfo DataInfo) {
	l.Data = append(l.Data, dataInfo)
}

func (l *LRUCache) Update(i int, dataInfo DataInfo) {
	l.Data[i] = dataInfo
	l.Data = append(l.Data[0:i], l.Data[i+1:]...)
	tmpData := []DataInfo{dataInfo}
	l.Data = append(tmpData, l.Data...)
	// l.DataMap[dataInfo.Key] = 0
}

func (l *LRUCache) IsFull() bool {
	if l.Len() == l.Cap {
		return true
	}
	return false
}

func (l *LRUCache) Remove() {

}

type DataInfo struct {
	Key     int
	Val     int
	UseTime int64
}

func Constructor(capacity int) LRUCache {
	list := make([]DataInfo, 0, capacity)
	// tmpMap := make(map[int]int, capacity)
	return LRUCache{
		Cap:  capacity,
		Data: list,
		// DataMap: tmpMap,
	}
}

func (this *LRUCache) Get(key int) int {
	res := -1
	// i, ok := this.DataMap[key]
	// if ok {
	// 	v := this.Data[i]
	// 	res = v.Val
	// 	v.UseTime = time.Now().UnixNano()
	// 	this.Update(i, v)
	// }
	for i, v := range this.Data {
		if v.Key == key {
			res = v.Val
			v.UseTime = time.Now().UnixNano()
			this.Update(i, v)
		}
	}
	// sort.Sort(this)
	return res
}

func (this *LRUCache) Put(key int, value int) {
	dataInfo := DataInfo{
		Key:     key,
		Val:     value,
		UseTime: time.Now().UnixNano(),
	}

	index, exist := 0, false
	// i, ok := this.DataMap[key]
	// if ok {
	// 	index = i
	// 	exist = true
	// }
	for i, v := range this.Data {
		if v.Key == key {
			index = i
			exist = true
		}
	}
	if !exist {
		if this.IsFull() {
			this.Data[this.Len()-1] = dataInfo
			this.Update(this.Len()-1, dataInfo)
		} else {
			this.Append(dataInfo)
			this.Data[this.Len()-1] = dataInfo
			this.Update(this.Len()-1, dataInfo)
		}
	} else {
		this.Update(index, dataInfo)
	}

	// sort.Sort(this)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	cache := Constructor(2)

	// cache.put(1, 1);
	// cache.put(2, 2);
	// cache.get(1);       // 返回  1
	// cache.put(3, 3);    // 该操作会使得关键字 2 作废
	// cache.get(2);       // 返回 -1 (未找到)
	// cache.put(4, 4);    // 该操作会使得关键字 1 作废
	// cache.get(1);       // 返回 -1 (未找到)
	// cache.get(3);       // 返回  3
	// cache.get(4);       // 返回  4
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(2, 1)
	fmt.Println(cache.Get(1))
	for _, v := range cache.Data {
		fmt.Println(v)
	}
	// arr := []int{1, 2, 3, 4, 5}
	// fmt.Println(arr)
	// arr = append(arr[0:4], arr[5:]...)
	// tmp := []int{7}
	// arr = append(tmp, arr...)
	// fmt.Println(arr)
}
