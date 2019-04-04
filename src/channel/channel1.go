package main

import "fmt"

//channel 理解为 队列
func channeldemo()  {

	c := make(chan int)

	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()

	c <- 1
	c <- 2

}

func main()  {

	channeldemo()

}