package main

import (
	"fmt"
	"imooc.com/mongo/learngo1/retriever/mock"
	"imooc.com/mongo/learngo1/retriever/real"
	"time"
)

const url  = "http://www.imooc.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, from map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func post(poster Poster)  {
	poster.Post(url,
		map[string]string {
			"name": "ccmouse",
			"course": "learngo",
		})
}
func download(r Retriever) string  {
	return r.Get(url)
}

func inspect(r Retriever)  {
	fmt.Println("Inspecting", r)
	fmt.Printf(">%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}

func main()  {
	var r Retriever
	retriever := mock.Retriever{"this is a fake imooc.com"}
	r = &retriever
	inspect(r)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	inspect(r)
	realRetriever, ok := r.(*real.Retriever)
	if ok {
		fmt.Println(realRetriever.TimeOut)
	} else {
		fmt.Println("not a real retriever")
	}

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
	//fmt.Println(download(r))
}
