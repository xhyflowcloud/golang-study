package main

import (
	"fmt"
)

func tryDefer() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error:" + err.Error())
		} else {
			panic(err)
		}
	}()
	//panic(errors.New("this is a error"))
	fmt.Println(5 / 0)
}
func main() {
	tryDefer()
}
