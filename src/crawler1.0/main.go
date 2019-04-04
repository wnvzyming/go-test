package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"io"
	"io/ioutil"
	"net/http"
)

func main()  {

	res,err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("error code" ,res.StatusCode  )
	}

	//i := DetermineEncoding(res.Body)

	//utf8reader := transform.NewReader(res.Body, i.NewDecoder())

	all , err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n",all)

}

func DetermineEncoding(r io.Reader) encoding.Encoding {

	bytes , err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e

}
