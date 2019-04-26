package main

import "fmt"

type MyCircularQueue struct {
	Val []int
	Head int
	Tail int
	Size int
}


/**初始化*/
func Constructor(k int) MyCircularQueue {
	obj := MyCircularQueue{
		Val: make([]int, k),
		Head: -1,
		Tail: -1,
		Size: k,
	}
	return obj
}


/**插入元素*/
func (this *MyCircularQueue) EnQueue(value int) bool {

	if this.IsFull() == true {
		return false;
	}
	if this.IsEmpty() == true {
		this.Head = 0;
	}
	this.Tail = (this.Tail + 1) % this.Size;
	this.Val[this.Tail] = value;

	return true;

}

/**删除元素*/
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() == true {
		return false
	}

	if this.Head == this.Tail {
		this.Head = -1
		this.Tail = -1
		return true
	}
	this.Head = (this.Head + 1) % this.Size

	return true

}

/**取头元素*/
func (this *MyCircularQueue) Front() int {

	if this.IsEmpty() == true {
		return -1
	}

	return this.Val[this.Head]

}


/**取尾元素*/
func (this *MyCircularQueue) Rear() int {

	if this.IsEmpty() == true {
		return -1
	}

	return this.Val[this.Tail]

}


/** 检查循环队列是否为空. */
func (this *MyCircularQueue) IsEmpty() bool {

	return this.Head == -1
}


/** 检查循环队列是否已满 */
func (this *MyCircularQueue) IsFull() bool {

	return (this.Tail + 1 ) % this.Size == this.Head

}

func main()  {

	fmt.Println(3%3)

	obj := Constructor(3);
	fmt.Println(obj.EnQueue(1))
	fmt.Println(obj.EnQueue(2))
	fmt.Println(obj.EnQueue(3))
	fmt.Println(obj.EnQueue(4))
	fmt.Println(obj.Rear())
	fmt.Println(obj.IsFull())

	//fmt.Println(obj)
	fmt.Println(obj.DeQueue())
	//fmt.Println(obj)

	fmt.Println(obj.EnQueue(4))
	fmt.Println(obj.Rear())

}
