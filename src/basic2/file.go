package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func readFile(filename string)  {

	file , err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scan := bufio.NewScanner(file)
	for scan.Scan(){
		fmt.Println(scan.Text())
	}

}

//死循环
func forever()  {
	for {
		fmt.Println("abc")
	}

}

func main()  {

	const file_name  = "src/basic2/abc.txt"
	content , err := ioutil.ReadFile(file_name)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("%s/n",content)
	}
	readFile(file_name)
	forever()

}
