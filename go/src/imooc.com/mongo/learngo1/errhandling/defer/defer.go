package main

import (
	"bufio"
	"fmt"
	"imooc.com/mongo/learngo1/funcational/flb"
	"os"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	defer fmt.Println(4)
}

func writeFile(filename string) {
	file, err := os.OpenFile(filename,
		os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := flb.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	writeFile("fib.txt")
}
