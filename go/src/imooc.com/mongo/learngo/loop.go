package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convert2binary(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lbs := n % 2
		result = strconv.Itoa(lbs) + result
	}
	return result
}

func printFile(filename string) {

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(
		convert2binary(5),
		convert2binary(13),
		convert2binary(1000),
	)
	printFile("abc.txt")
}
