package main

import (
	"fmt"
)

func for_arr( arr [3]int)  {
	arr[0] = 100
	for i := range arr  {
		fmt.Println(arr[i])
	}
}

func maxProfit(prices []int) int {
	n := 0
	for i:=1; i<len(prices);i++ {
		if prices[i] > prices[i-1]{

			n += prices[i] - prices[i-1]
		}
	}


	return n

}


func for_arr_zhizhen( arr *[3]int)  {
	arr[0] = 100
	for i := range arr  {
		fmt.Println(arr[i])
	}
}

func main()  {

	r := []int{7,1,5,3,6,4}

	fmt.Print(maxProfit(r))

	//数组是值类型 值类型是拷贝，不会改变原数组的值 引用类型 会改变
	//arr1 :=  [3]int{1,2,3}
	//arr2 :=  [3]int{99,98,93}
	//
	//fmt.Println("值传递数组")
	//for_arr(arr1)
	//
	//fmt.Println("引用传递传递数组")
	//for_arr_zhizhen(&arr2)
	//
	//fmt.Println("原数组")
	//for i := range arr1  {
	//	fmt.Println(arr1[i])
	//}
	//
	//for i := range arr2  {
	//	fmt.Println(arr2[i])
	//}

	//beego.Run("127.0.0.1:80891")
}
