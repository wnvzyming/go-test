package main

import (
	"fmt"
	"sync"
	"time"
)

type auto struct {
	value int
	lock sync.Mutex
}

func (a *auto)  add() {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value++
}

func (a *auto)  get() int{
	a.lock.Lock()
	defer a.lock.Unlock()
	return int(a.value)
}

func main()  {

	var a auto

	go func() {
		a.add()
	}()

	time.Sleep(time.Millisecond)

	fmt.Println(a.value)

	//fmt.Println(a.get())

}
