package main

import "fmt"

func main() {
	bs1 := []byte("hello world!")
	bs2 := make([]byte, 5)
	copy(bs2, bs1)
	fmt.Println(string(bs2))
	bs3 := make([]byte, 5, 100)
	copy(bs3, bs1)
	fmt.Println(string(bs3))
}






