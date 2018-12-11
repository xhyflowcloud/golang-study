package main

import (
	"fmt"
	"imooc.com/mongo/learngo1/retriever/mock"
	"imooc.com/mongo/learngo1/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string  {
	return r.Get("http://www.imooc.com")
}

func main()  {
	var r Retriever
	r = mock.Retriever{"this is a fake imooc.com"}
	r = real.Retriever{}
	fmt.Println(download(r))
}
