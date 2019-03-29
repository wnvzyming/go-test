package main

import "fmt"

func main()  {
	s := "奸雄曹操xixi!"

	for _ , v := range []byte(s) {
		//utf8编码 可变字节，英文1个字节 汉子占3个字节
		//E5 A5 B8 E9 9B 84 E6 9B B9 E6 93 8D 78 69 78 69 21
		fmt.Printf("%X ",v)
	}

	fmt.Println("go 字符串处理，转为 rune 类型的slice处理")

	for rk, rv := range []rune(s) {
		fmt.Printf("%d , %c ",rk, rv)

	}
}
