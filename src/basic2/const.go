package main

import "fmt"

func emus()  {
	const (
		java = iota
		php
		golang
	)

	fmt.Println(java,php,golang)
}

func return_num( num int )  int{

	if num > 100 {
		return 100
	} else if num < 100 {
		return 99
	}
	return num
}

func main()  {

	emus();

	num := return_num(200)
	fmt.Println(num)

}
