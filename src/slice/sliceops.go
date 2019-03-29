package main

import "fmt"

func check_cap( arr []int)  {

	fmt.Println("arr len %d , arr cap %d" , len(arr), cap(arr))

}

func main()  {

	var arr []int

	s1 := make([]int,10,32)

	for i:=0 ; i<100 ; i++  {
		check_cap(arr)
		arr = append(arr,i * 2 +1)
	}

	fmt.Println(arr)

}
