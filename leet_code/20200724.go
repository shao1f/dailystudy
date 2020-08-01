package main

import "fmt"

type MyCircularQueue struct {
	Data []int
	Len  int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor1(k int) MyCircularQueue {
	return MyCircularQueue{
		Data: make([]int, 0, k),
		Len:  k,
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if len(this.Data) == this.Len {
		return false
	}
	this.Data = append(this.Data, value)
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if len(this.Data) == 0 {
		return false
	}
	this.Data = this.Data[:len(this.Data)-1]
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if len(this.Data) == 0 {
		return -1
	}
	return this.Data[0]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if len(this.Data) == 0 {
		return -1
	}
	return this.Data[len(this.Data)-1]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return len(this.Data) == 0
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return len(this.Data) == this.Len
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */

func main() {
	q := Constructor1(8)
	fmt.Println(q.EnQueue(3))
	fmt.Println(q.EnQueue(9))
	fmt.Println(q.EnQueue(5))
	fmt.Println(q.EnQueue(0))
	fmt.Print(q.DeQueue())
	fmt.Print(q.DeQueue())
	fmt.Print(q.IsEmpty())
	fmt.Print(q.Rear())
	fmt.Print(q.Rear())
}
