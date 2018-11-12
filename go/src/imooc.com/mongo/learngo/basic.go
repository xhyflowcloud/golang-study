package main

import "fmt"

var (
	aa = 1
	bb = "hello"
	ss = true
	tt = 3
)

func variableDef() {
	a, b, c, d := 1, "ok", true, "ddd"
	fmt.Println(a, b, c, d)
}
func main() {
	fmt.Println("Hello World")
	fmt.Println(aa, bb, ss)
	fmt.Println(aa+tt, bb, ss)
	variableDef()
}
