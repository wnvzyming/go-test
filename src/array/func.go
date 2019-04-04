package main

type intfunc func() int

func (g intfunc) Read(p []byte) (num int , err error)  {

	next := g()
	
}


