package main

import "fmt"

func score( score int)  {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic("score is wrong %d")
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "b"
	case score <= 100:
		g = "a"
	}
	fmt.Print(g)

}

func sum(numbers ...int) int {

	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}

func main()  {
	fmt.Print(sum(1,2,3,4,5,5))
	score(99)
}
