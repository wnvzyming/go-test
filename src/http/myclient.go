package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main()  {

	s , err := http.Get("https://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer s.Body.Close()

	c,err := httputil.DumpResponse(s ,true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n",c)


}
