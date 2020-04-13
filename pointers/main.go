package main

import "fmt"

func main() {
	var i int = 11

	// nil memory
	// var ptr *int

	// -- this stmt gives memory to the ptr
	var ptr *int = new(int)

	fmt.Println(i, ptr)

	ptr = &i
	fmt.Println(ptr, *ptr, i)

	// ptr arithmetic not allowed by gopher
	// fmt.Println(*(ptr + 1))

}
