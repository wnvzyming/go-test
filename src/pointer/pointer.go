package main

import "fmt"

type demo struct {
	Value string
}



func reverse(p *[]int) {
	for i, j := 0, len(*p)-1; i < len(*p)/2; i, j = i+1, j-1 {
		(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
	}
}

func main() {
	v := []int{1, 2, 3, 4}
	reverse(&v)
	fmt.Println(v)

	var pt *demo //*用作指针类型声明
	pt = &demo{Value: "val111"} //&取值地址
	d1 := *pt //*用作取地址指向的实际值
	fmt.Println(d1)

}