package main

import "fmt"

/**
执行用时 : 8 ms, 在Add Digits的Go提交中击败了96.55% 的用户
内存消耗 : 2.2 MB, 在Add Digits的Go提交中击败了59.26% 的用户
 */

func addDigits( num int)  int {

	if num > 9 {
		num = num % 9;
		if num == 0 {
			return 9
		}
	}

	return num
}

func main()  {

	fmt.Print(addDigits(38))
	
}


