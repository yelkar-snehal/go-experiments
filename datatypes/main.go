package main

import "fmt"

func main() {
	// both of the following are explicitly declared
	var i int
	i = 11
	fmt.Println("i int", i)

	var f float32 = 3.142
	fmt.Println("f float32", f)

	// implicit
	str := "Snehal"
	fmt.Println("str string implicit initialization", str)

	b := true
	fmt.Println("b boolean", b)

	c := complex(3, 4)
	fmt.Println("c complex", c)
	fmt.Println("real and img of c", real(c), imag(c))

	// const has to be explicit
	const pi float32 = 3.14
	// not allowed duh!
	// pi = 1.3
	fmt.Println("pi const", pi)
}
