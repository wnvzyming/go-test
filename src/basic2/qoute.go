package main

import "fmt"

func val( a int) int {
	return a
}
//
//func  qoute( a *int)  int {
//	return a++
//}

func main()  {
	a := 3
	fmt.Print(val(a))
}