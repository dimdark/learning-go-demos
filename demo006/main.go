package main

import "fmt"

type MyInt int

func main() {
	var a MyInt = 1
	fmt.Println(a)

	b := new(MyInt)
	*b = 2
	fmt.Println(*b)
}
