package main

import "fmt"

/**
map  key
   	map 使用哈希表，必须可以比较相等
	除了slice map function的内建类型，都可以当key
	struct 类型不包括上述字段，也可以做key
 */

func sub_str( s string)  {

	last_ch := make(map[byte]int)
	start := 0
	maxLength := 0

	for i,ch := range []byte(s) {

		fmt.Println(i,ch)

		lastI , ok := last_ch[ch]

		if ok && lastI >= start {
			start = lastI + 1
		}

		if i-start+1 >maxLength {
			maxLength = i-start +1
		}

		last_ch[ch] = i
	}

	fmt.Println(maxLength)
}


func main()  {

	//fmt.Println(sub_str("a"))
	//fmt.Println(sub_str("abcahgjkkio"))
	sub_str("abcahgjkkio")

	m := map[string]string{
		"name" : "lb",
		"age" : "888",
		"sex" : "1",
	}
	m1 := make(map[string] int)
	var m2 map[string] string

	fmt.Println(m)
	fmt.Println(m1)
	fmt.Println(m2)

}
