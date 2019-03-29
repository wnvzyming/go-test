package main

import "fmt"

func main()  {

	//slice 是数组底层的一个view
	arr := [...]int {0,1,2,3,4,5,6,7}

	s1 := arr[2:6]
	fmt.Println(s1) // [2,3,4,5]
	s2 := s1[3:5]
	fmt.Println(s2) // [5,6] slice 可以向后扩展，不可以向前扩展 比如 取 [1,3] 就不行

}
