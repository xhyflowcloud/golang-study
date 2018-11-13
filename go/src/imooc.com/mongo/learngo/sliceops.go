package main

import "fmt"

func printSlice(s []int) {
	fmt.Println("s = %v, len(s) = %d, cap(s) = %d", s, len(s), cap(s))
}
func main() {
	var s []int
	for i := 0; i < 100; i++ {
		//printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{1, 2, 3, 4, 5}
	printSlice(s1)

	s2 := make([]int, 10)
	s3 := make([]int, 10, 16)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(front, tail)
}
