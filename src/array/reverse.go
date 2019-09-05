package main

import (
	"fmt"
)

//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

func reverse(x int ) int {
	rev := 0
	for x != 0 {
		//if rev < math.MinInt64/10 {
		//	return 0
		//}
		//if rev > math.MaxInt64/10 {
		//	return 0
		//}

		v := x % 10
		x /= 10
		rev = rev * 10 + v
		//if rev > math.MaxInt64 || rev < math.MinInt64 {
		//	return 0
		//}

	}





	return rev

}

func main()  {

	//fmt.Print(1 % 10)

	fmt.Print(reverse(-123))
}