package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is ccmouse@gmail.com
email1 is abc@def.org
email2 is    kkk@qq.com
`

func main() {
	re, err := regexp.Compile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)\.([a-zA-Z0-9]+)`)
	if err != nil {
		panic(err)
	}
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		for _, e := range m {
			fmt.Printf("%s ", e)
		}
		fmt.Println()
	}
	//fmt.Println(match)

}
